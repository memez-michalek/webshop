package helpers

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
)

func getAuthCredentials(c *gin.Context) (string, string, error) {

	username, err := c.Cookie("username")
	if err != nil || username == "" {
		log.Print(err)
		return "", "", errors.New("username cookie is empty")
	}
	email, err := c.Cookie("email")
	if err != nil || email == "" {
		log.Print(err)
		return "", "", errors.New("email cookie is empty")
	}

	return username, email, nil

}
func IsAuthenticated(c *gin.Context) bool {
	_, _, err := getAuthCredentials(c)
	if err != nil {

		return false
	}
	return true
}
func ErrorChecker(e error) {
	if e != nil {
		log.Print(e)
	}
}

func IsValid(token string) error{
	
	return nil

}


