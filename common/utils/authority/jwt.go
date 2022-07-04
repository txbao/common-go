package authority

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
//使用说明
sdk := authority.NewSdk("1111", 60)
	dataCliams := authority.DataClaims{
		UserId:   1,
		UserName: "kefu2",
	}
	token, expiresAt, err := sdk.MakeCliamsToken(dataCliams)
	fmt.Println("token:", token, err)

	orgToken, err := sdk.ParseCliamsToken(token)
	if err == nil {
		fmt.Println(orgToken.Data.UserName)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println(expiresAt)
	return
*/

type DataClaims struct {
	UserId     uint      `json:"user_id"`
	UserName   string    `json:"user_name"`
	CreateTime time.Time `json:"create_time"`
}
type TokenClaims struct {
	Data DataClaims `json:"data"`
	jwt.StandardClaims
}

type jwtStruct struct {
	AccessSecret string `json:"access_secret"`
	AccessExpire int64  `json:"access_expire"`
}

func NewSdk(secret string, expireSecond int64) *jwtStruct {
	return &jwtStruct{
		AccessSecret: secret,
		AccessExpire: expireSecond,
	}
}

//生成jwt
func (o *jwtStruct) MakeCliamsToken(data DataClaims) (string, int64, error) {
	ExpiresAt := time.Now().Unix() + o.AccessExpire
	tokenCliams := TokenClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenCliams)
	tokenString, err := token.SignedString([]byte(o.AccessSecret))
	return tokenString, ExpiresAt, err
}

//解析jwt
func (o *jwtStruct) ParseCliamsToken(token string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(o.AccessSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
