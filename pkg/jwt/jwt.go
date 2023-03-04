/**
 * @Author: 戈亓
 * @Description:
 * @File:  jwt
 * @Version: 1.0.0
 * @Date: 2022/7/24 17:55
 */

package jwt

import (
	"errors"
	"time"
)
import "github.com/dgrijalva/jwt-go"

type TokenData struct {
	UserId     string
	Realname   string
	UserType   string
	Unit       string
	Department string
	jwt.StandardClaims
}

const TokenExperieDuration = time.Hour * 2

var MySecret = []byte("thisislaoliusecret")

func (c *TokenData) GenToken(userid, realname, usertype, unit, department string) (string, error) {
	c.UserId = userid
	c.Realname = realname
	c.Department = department
	c.UserType = usertype
	c.Unit = unit
	c.ExpiresAt = time.Now().Add(TokenExperieDuration).Unix()
	c.Issuer = "cea_api"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(MySecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*TokenData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenData{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*TokenData); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
