package errorrpc

//用户服务code
const USER_ERROR = 11000

var (
	UserErr = NewCodeError(USER_ERROR+1, "用户错误")
)
