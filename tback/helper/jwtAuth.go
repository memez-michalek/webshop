package helper

import (
	"errors"
	"log"
	"tback/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(email string, password string) (string, error) {
	if email == "" || password == "" {
		log.Print("values are empty")
		return "", errors.New("values provided are empty")
	}
	data := jwt.NewWithClaims(jwt.SigningMethodHS256, model.Auth{
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Hour * 8)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := data.SignedString(JWTSECRET)
	if err != nil {

		log.Print("could not create token", err)
	}
	return token, nil
}

func GetTokenValue(token string) (string, string, error) {
	secret, err := jwt.ParseWithClaims(token, new(model.Auth), func(token *jwt.Token) (interface{}, error) {
		return JWTSECRET, nil
	})
	if err != nil {
		log.Print("could not load jwt secret: ", err)
	}
	if claim, err := secret.Claims.(*model.Auth); err {
		if claim.ExpiresAt < time.Now().Unix() {
			return "", "", errors.New("token expired")
		}
		return claim.Email, claim.Password, nil
	}
	return "", "", errors.New("could not parse claim ")

}
