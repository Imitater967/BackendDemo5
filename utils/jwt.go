package utils

import (
	"ByteTechTraining/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	ID int64
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var MyKey = []byte("secret") // 密钥 后续通过统一的配置文件读取

func GenerateToken(model models.UserAuthModel) (string, error, time.Time) {
	exp := time.Now().Add(TokenExpireDuration)
	myClaims := MyClaims{
		ID: model.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp.Unix(), // 过期时间
			Issuer:    "Frogs",    // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	tokenString, err := token.SignedString(MyKey)
	if err != nil {
		return "", err, time.Time{}
	}
	return tokenString, nil, exp
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MyKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
