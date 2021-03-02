package webtoken

import (
	"api/errorCodes"
	"api/models"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRETTOKEN = []byte("fjklajfalfjop4pr90409fj03rj030f30")

func CreateToken(username string, email string, apikey string) (string, error) {
	if username == "" || email == "" || apikey == "" {
		return "", errors.New(errorCodes.TOKENERROR)
	}
	data := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JWTAUTH{
		username,
		email,
		apikey,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Hour * 10)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := data.SignedString(SECRETTOKEN)
	if err != nil {
		log.Print(err)
	}

	return token, nil

}
func GetValidTokenValue(token string) (string, string, string, error) {
	log.Print(token)
	secret, err := jwt.ParseWithClaims(token, new(models.JWTAUTH), func(token *jwt.Token) (interface{}, error) {
		return SECRETTOKEN, nil
	})
	log.Print(secret)
	if err != nil {
		log.Print(err)
		log.Print("Token related error")
		return "", "", "", errors.New(errorCodes.TOKENERROR)
	}
	if claim, err := secret.Claims.(*models.JWTAUTH); err {
		if claim.ExpiresAt < time.Now().Unix() {
			return "", "", "", errors.New(errorCodes.TOKENEXPIRED)
		}
		return claim.Username, claim.Email, claim.Apikey, nil
	}
	return "", "", "", errors.New(errorCodes.CLAIMERROR)
}

func GetInvalidTokenValue(token string) (string, string, string) {
	secret, err := jwt.ParseWithClaims(token, new(models.JWTAUTH), func(token *jwt.Token) (interface{}, error) {
		return SECRETTOKEN, nil
	})
	if err != nil {
		log.Print("Token related error")
		return "", "", ""
	}
	if claim, err := secret.Claims.(*models.JWTAUTH); err {

		return claim.Username, claim.Email, claim.Apikey
	}
	return "", "", ""

}
