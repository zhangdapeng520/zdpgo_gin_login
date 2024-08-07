package zdpgo_gin_login

import (
	"errors"
	jwt "github.com/zhangdapeng520/zdpgo_jwt"
	"time"
)

type MyClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetToken 生成Token
func GetToken(id int, username, jwtKey string) (string, error) {
	c := MyClaims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Second * 1800)},
			Issuer:    username,
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString([]byte(jwtKey))
}

// ParseToken 解析Token
func ParseToken(tokenString, jwtKey string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("无效的Token")
}
