package logic

import (
    "net/http"
	{{.imports}}
)

type {{.logic}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) {{.logic}} {
	return {{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) (interface{}, error) {
	// todo: add your logic here and delete this line

    //数据验证
    /*
	formStruct := &valid.Index{
		ActivityId: req.ActivityId,
	}
	if err := validators.Valid(formStruct); err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	*/
	/*
	//获取用户信息
    tokenSdk := authority.NewToken(l.svcCtx.Rds, l.svcCtx.Config.Auth.AccessExpire)
    authorization := l.r.Header.Get(authority.AuthorizationTag)
    tokenData, err := tokenSdk.GetTokenData(authorization)
    if err != nil {
        return nil, errorx.NewDefaultError(err.Error())
    }
    userId := tokenData.UserId
	*/


	return nil, nil
}
