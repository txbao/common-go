package {{.pkgName}}

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

func New{{.logic}}(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *{{.logic}} {
	return &{{.logic}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *{{.logic}}) {{.function}}({{.request}}) {{.responseType}} {
	// todo: add your logic here and delete this line
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

	{{.returnString}}
}
