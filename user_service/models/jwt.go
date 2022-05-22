package models

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTWrapper struct{
	SecretKey string
	Issuer string
	ExpirationTime int64
}

type JWTClaims struct{
	Email string
	jwt.StandardClaims
}

func (j *JWTWrapper) GenerateToken(email string) (signedToken string, err error){
	claims := JWTClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(j.ExpirationTime)).Unix(),
			Issuer: j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	return 
}

func (j *JWTWrapper) ValidateToken(signedToken string) (claims *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", t.Header["alg"]))
		}

		key, err := base64.URLEncoding.DecodeString(j.SecretKey)
		if err != nil{
			return nil, errors.New(fmt.Sprintf("Unable to decode secret key: %x", err))
		}

		return key, nil
	})

	if err != nil{
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); token.Valid && ok{
		if claims.ExpiresAt < time.Now().Local().Unix(){
			err = errors.New("JWT is expired")
		}
	}
	
	return
}