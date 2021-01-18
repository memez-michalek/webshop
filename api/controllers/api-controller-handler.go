package controllers

import (
	"api/database"
	"api/models"
	"api/webtoken"
	"log"

	"github.com/gin-gonic/gin"
)

func LogInToApi(c *gin.Context) {
	var apiuser models.APILOGIN

	err := c.ShouldBindJSON(&apiuser)
	log.Print(err)
	if err != nil {
		c.JSON(400, "could not bind request with model")

	} else {

		err := database.ApiLogin(apiuser.Email, apiuser.Username, apiuser.Key)
		if err != nil {
			c.JSON(400, "wrong credentials: "+err.Error())
		} else {
			token := webtoken.CreateToken(apiuser.Username, apiuser.Email, apiuser.Key)
			c.JSON(200, token)
		}

	}

}



