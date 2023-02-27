package impl

import (
	"errors"
	"fmt"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtSecurityServiceImpl struct {
}

func (j JwtSecurityServiceImpl) GenerateJWTToken(email string) (string, error) {
	var sampleKey = []byte("MYSECRETEKEY")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(2 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = email

	tokenString, err := token.SignedString(sampleKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j JwtSecurityServiceImpl) VerifyToken(tokenString string) error {
	var sampleKey = []byte("MYSECRETEKEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there is an error in parsing")
		}
		return sampleKey, nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if token == nil {

	}
	return nil
}

func (j JwtSecurityServiceImpl) GetEmail(tokenString string) (string, error) {

	var sampleKey = []byte("MYSECRETEKEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there is an error in parsing")
		}
		return sampleKey, nil
	})

	if err != nil || token == nil {
		fmt.Println(err)
		return "", nil
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("")
	}
	email := claims["user"].(string)
	if email == "" {
		return "", errors.New("")
	}
	return email, nil
}

func (j JwtSecurityServiceImpl) IsValid(tokenString string) bool {
	var sampleKey = []byte("MYSECRETEKEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there is an error in parsing")
		}
		return sampleKey, nil
	})
	if err != nil || token == nil {
		fmt.Println(err)
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		return false
	}
	return true
}

func (j JwtSecurityServiceImpl) GenerateSecretKey() string {
	token := jwt.New(jwt.SigningMethodEdDSA)
	var sampleKey = []byte("MYSECRETEKEY")
	tokenString, err := token.SignedString(sampleKey)
	if err != nil {
		return ""
	}
	return tokenString
}

func NewJwtSecurityService() *JwtSecurityServiceImpl {
	return &JwtSecurityServiceImpl{}
}

var _ service.JwtSecurityService = &JwtSecurityServiceImpl{}
