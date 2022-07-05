package utils

import (
	"fmt"
	"google.golang.org/grpc/status"
	"net/http"
	"regexp"
)

//判断
func RegIsMatch(str string, expr string) bool {
	reg, err := regexp.Compile(expr)
	if err != nil {
		fmt.Println(err)
	}
	return reg.MatchString(str)
}

//判断
func RegIsMatchMust(str string, pattern string) bool {
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}

//判断是否微信浏览器
func RegIsMicroMessenger(r *http.Request) bool {
	ua2 := r.Header.Get("User-Agent")
	reg := regexp.MustCompile("MicroMessenger")

	return reg.MatchString(ua2)
}

//判断是否雀省APP里
func RegIsQueshenApp(r *http.Request) bool {
	ua2 := r.Header.Get("User-Agent")
	reg := regexp.MustCompile("qssq_app")

	return reg.MatchString(ua2)
}

//判断是否在小程序内
func RegIsMiniprogram(r *http.Request) bool {
	ua2 := r.Header.Get("User-Agent")
	reg := regexp.MustCompile("miniProgram")

	return reg.MatchString(ua2)
}

//判断是否是iOS
func RegIsIOS(r *http.Request) bool {
	ua2 := r.Header.Get("User-Agent")
	reg := regexp.MustCompile("Mac OS")

	return reg.MatchString(ua2)
}

//判断是否是上下文超时
func RegIsIDeadlineExceeded(err error) bool {
	s, _ := status.FromError(err)
	if s.Code().String() == "DeadlineExceeded" {
		return true
	}
	return false
}
