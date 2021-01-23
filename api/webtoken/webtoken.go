package webtoken

import (
	"api/models"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRETTOKEN = []byte("he0ErCNloe4J7Id0Ry2SEDg09lKkZkfsRiGsdX_vgEg")

func CreateToken(username string, email string, apikey string) string {

	data := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JWTAUTH{
		username,
		email,
		apikey,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(time.Minute * 10)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := data.SignedString(SECRETTOKEN)
	if err != nil {
		log.Print(err)
	}

	return token

}
func GetTokenValue(token string) (string, string, string, error) {
	secret, err := jwt.ParseWithClaims(token, new(models.JWTAUTH), func(token *jwt.Token) (interface{}, error) {
		return SECRETTOKEN, nil
	})
	if err != nil {
		log.Print("Token related error")
		return "", "", "", errors.New("token error")
	}
	if claim, err := secret.Claims.(*models.JWTAUTH); err {
		if claim.ExpiresAt < time.Now().Unix() {
			return "", "", "", errors.New("token expired")
		}
		return claim.Username, claim.Email, claim.Apikey, nil
	}
	return "", "", "", errors.New("claim error")
}
