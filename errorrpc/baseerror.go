package errorrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/txbao/common-go/proto"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
	"runtime/debug"
)

type CodeError struct {
	Code    int64    `json:"code"`
	Msg     string   `json:"msg"`     //给API的显示信息
	ErrMsg  string   `json:"err_msg"` //用于日志的错误信息
	Details []string `json:"details,omitempty"`
	Stack   []string `json:"stack,omitempty"`
}

type RpcError interface {
	GRPCStatus() *status.Status
}

func (c CodeError) ToRpc() error {
	s, _ := status.New(codes.Unknown, c.Msg).
		WithDetails(&proto.Error{Code: int32(c.Code), Message: c.Msg, Detail: c.Details})
	return s.Err()
}

func ToRpc(err error) error {
	if err == nil {
		return err
	}
	switch err.(type) {
	case RpcError:
		return err
	case *CodeError:
		return err.(*CodeError).ToRpc()
	default:
		return Fmt(err).ToRpc()
	}
}

func (c CodeError) WithMsg(msg string) *CodeError {
	return &CodeError{Code: c.Code, Msg: msg}
}
func (c CodeError) WithErrMsg(errMsg string) *CodeError {
	msg := c.Msg
	if errMsg == "sql: no rows in result set" {
		msg = "没有记录"
	}
	return &CodeError{Code: c.Code, Msg: msg, ErrMsg: errMsg}
}
func (c CodeError) AddDetail(msg ...interface{}) *CodeError {
	c.Details = append(c.Details, fmt.Sprint(msg))
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	//返回数据太大了，屏蔽
	logx.Error("ErrStack:"," msg:",fmt.Sprint(msg)," Stack:",string(debug.Stack()))
	c.Stack = append(c.Stack, string(debug.Stack()))
	return &c
}

func (c CodeError) AddDetailf(format string, a ...interface{}) *CodeError {
	c.Details = append(c.Details, fmt.Sprintf(format, a...))
	return &c
}

func (c *CodeError) GetDetailMsg() string {
	if len(c.Details) == 0 {
		return c.Msg
	}
	return fmt.Sprintf("msg=%s,detail=%v", c.Msg, c.Details)
}

func NewCodeError(code int64, msg string) *CodeError {
	return &CodeError{Code: code, Msg: msg, ErrMsg: msg}
}

func NewDefaultError(msg string) error {
	return Default.WithMsg(msg)
}

func (e CodeError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

//将普通的error及转换成json的error或error类型的转回自己的error
func Fmt(errs error) *CodeError {
	if errs == nil {
		return nil
	}
	switch errs.(type) {
	case *CodeError:
		return errs.(*CodeError)
	case RpcError: //如果是grpc类型的错误
		s, _ := status.FromError(errs)
		if s.Code() != codes.Unknown { //只有自定义的错误,grpc会返回unknown错误码
			err := fmt.Sprintf("rpc err detail is nil|err=%#v", s)
			return System.AddDetail(err)
		}
		var ret CodeError
		err := json.Unmarshal([]byte(s.Message()), &ret)
		if err != nil {
			return System.AddDetail(err)
		}
		return &ret
	default:
		var ce CodeError
		err := json.Unmarshal([]byte(errs.Error()), &ce)
		if err != nil {
			return System.AddDetail(errs.Error())
		}
		return Default.AddDetail(errs)
	}
}

func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		logx.WithContext(ctx).Errorf("err=%s", Fmt(err).Error())
	} else {
		logx.WithContext(ctx).Slowf("resp=%+v", resp)
	}
	//err = ToRpc(err)
	return resp, err
}

func Cmp(err1 error, err2 error) bool {
	if err2 == nil && err1 == nil {
		return true
	}
	if err1 == nil || err2 == nil {
		return false
	}
	return Fmt(err1).Code == Fmt(err2).Code
}
