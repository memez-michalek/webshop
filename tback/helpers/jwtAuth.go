package helpers

import (
	"log"
	"templateBackend/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func CreateToken(email string, password string) (string, error) {
	if email == "" || password == "" {
		log.Print("values are empty")
		return "", errors.Error("values provided are empty")
	}
	data := jwt.NewWithClaims(jwt.SigningMethodHS256, models.Auth{
		Email:    email,
		Password: password,
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

func GetTokenValue(token string) (string, string, err) {
	secret, err := jwt.ParseWithClaims(token, new(models.Auth), func(token *jwt.Token) (interface{}, error) {
		return JWTSECRET, nil
	})
	if err != nil {
		log.Print("could not load jwt secret: ", err)
	}
	if claim, err := secret.Claims.(*models.Auth); err {
		if claim.ExpiresAt < time.Now().Unix() {
			return "", "", errors.Error("token expired")
		}
		return claim.Email, claim.Password, nil
	}
	return "", "", errors.Error("could not parse claim ", err)

}
