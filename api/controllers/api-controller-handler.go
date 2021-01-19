package controllers

import (
	"api/database"
	"api/models"
	"api/webtoken"
	"log"

	"github.com/gin-gonic/gin"
)

func LogInToApi(c *gin.Context) {
	apiuser := new(models.APILOGIN)

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
			mainSiteContent := database.Addapi("djdjjd")
			c.JSON(200, token)
			
		}
	}
}
func GetProduct(c *gin.Context) {
	product := new(models.Product)
	err := c.ShouldBind(&product)
	if err != nil{
		c.JSON(400, "could not bind " + err.Error())
	}
	//err := database.GetProducts(product.)

}
