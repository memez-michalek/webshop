package controllers

import (
	"api/database"
	"api/helpers"

	"log"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	if helpers.IsAuthenticated(c) {
		c.Redirect(302, "/")
	}
	email := c.PostFormArray("email")[0]
	username := c.PostFormArray("username")[0]
	password := c.PostFormArray("password")[0]
	loggedin := database.LoginUser(email, username, password)
	if !loggedin {
		log.Print("not logged in")
		c.Redirect(302, "/login")
	} else {

		c.SetCookie("username", username, 3600, "/", "localhost", false, true)
		c.SetCookie("email", email, 3600, "/", "localhost", false, true)
		c.Redirect(302, "/")

	}

}
func RegisterHandler(c *gin.Context) {
	if helpers.IsAuthenticated(c) {

		c.Redirect(302, "/")
	} else {
		email := c.PostFormArray("email")[0]
		username := c.PostFormArray("username")[0]
		password := c.PostFormArray("password")[0]
		inserted := database.InsertUser(email, username, password)
		if !inserted {
			log.Print("not inserted")
			c.Redirect(302, "/register")
		}
		c.SetCookie("username", username, 3600, "/", "localhost", false, true)
		c.SetCookie("email", email, 3600, "/", "localhost", false, true)
		c.Redirect(302, "/")

	}
}
func LogoutHandler(c *gin.Context) {
	if helpers.IsAuthenticated(c) {
		c.SetCookie("username", "", 3600, "/", "localhost", false, true)
		c.SetCookie("email", "", 3600, "/", "localhost", false, true)
		c.Redirect(302, "/")
	} else {
		c.Redirect(302, "/")
	}
}
func GetApiHandler(c *gin.Context) {
	if helpers.IsAuthenticated(c) {
		shopname := c.PostFormArray("name")[0]
		email, err := c.Cookie("email")
		helpers.ErrorChecker(err)
		username, err := c.Cookie("username")
		helpers.ErrorChecker(err)
		added := database.AddApi(email, username, shopname)
		if added {
			log.Print("api created")
			c.Redirect(302, "/")
		} else {
			log.Print("api has not been created")
			c.Redirect(302, "/")
		}
	} else {
		c.Redirect(302, "/login")
	}

}
