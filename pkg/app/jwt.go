package app

import (
	"time"

	"BlogService/global"
	"BlogService/pkg/errcode"
	"BlogService/pkg/util"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.RegisteredClaims
}

// GetJWTSecret 获取服务器私钥
func GetJWTSecret() []byte {
	return []byte(global.Config.JWT.Secret)
}

// GenerateToken 生成并签名Token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.Config.JWT.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expireTime},
			NotBefore: &jwt.NumericDate{Time: nowTime},
			IssuedAt:  &jwt.NumericDate{Time: nowTime},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetJWTSecret())
	return tokenString, errors.Wrap(err, "JWT signing failed")
}

// ParseToken 解析并校验Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil { //解析失败
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return nil, errcode.UnauthorizedTokenTimeout
		default:
			return nil, errcode.UnauthorizedTokenError
		}
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid { //校验成功
		return claims, nil
	}
	return nil, errcode.UnauthorizedTokenError
}
