package controllers

import (
	"api/database"
	"api/models"
	"api/webtoken"
	"log"
	"api/errorCodes"
	"github.com/gin-gonic/gin"
)

func LogInToApi(c *gin.Context) {
	var (
		apiuser = new(models.APILOGIN)
		shop    = new(models.Shop)
	)
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
			queriedShop, err := database.QueryShopByApiKey(models.SHOPLIST, apiuser.Key)
			if err != nil {
				shop.ApiKey = apiuser.Key
				shop.Name = queriedShop.Name
				models.SHOPLIST = append(models.SHOPLIST, *shop)
				c.JSON(200, token)
			} else {
				log.Print(err)
				c.JSON(400, "shop has already been logged in")
				c.JSON(400, err)
			}
		}
	}
}
func MainPage(c *gin.Context) {
	apiUSER := new(models.APIUSER)
	err := c.ShouldBind(&apiUSER)
	if err != nil {
		c.JSON(400, "could not bind to model"+err.Error())
	} else {
		username, email, apikey, err := webtoken.GetTokenValue(apiUSER.Token)
		if err != nil{
			switch{
				case err.Error() == ""
			}
		}
	}
}
