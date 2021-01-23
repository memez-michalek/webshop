package views

import (
	"api/database"
	"api/helpers"
	"log"

	"github.com/gin-gonic/gin"
)

func IndexView(c *gin.Context) {
	if !helpers.IsAuthenticated(c) {
		c.HTML(200, "index.html", gin.H{
			"Username": "",
		})
	} else {
		username, err := c.Cookie("username")
		helpers.ErrorChecker(err)
		email, err := c.Cookie("email")
		helpers.ErrorChecker(err)
		apikeys, err := database.GetApiKeys(email, username)
		helpers.ErrorChecker(err)
		c.HTML(200, "index.html", gin.H{
			"Username": username,
			"Apikeys":  apikeys,
		})
	}
}
func LoginView(c *gin.Context) {
	if !helpers.IsAuthenticated(c) {
		c.HTML(200, "register-login.html", "")
	} else {
		c.Redirect(302, "/")
	}
}
func RegisterView(c *gin.Context) {
	if !helpers.IsAuthenticated(c) {
		c.HTML(200, "register-login.html", "")
	} else {
		c.Redirect(302, "/")
	}
}
func GetApiView(c *gin.Context) {
	if helpers.IsAuthenticated(c) {
		c.HTML(200, "getapi.html", "")
	} else {
		log.Print("not authenticated")
		c.Redirect(302, "/")
	}
}
