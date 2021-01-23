package controllers

import (
	"api/database"
	"api/errorCodes"
	"api/models"
	"api/webtoken"
	"errors"
	"log"

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
		c.JSON(400, errorCodes.COULDNOTBIND)

	} else {

		err := database.ApiLogin(apiuser.Email, apiuser.Username, apiuser.Key)
		if err != nil {
			c.JSON(400, "wrong credentials: "+err.Error())
		} else {
			token, err := webtoken.CreateToken(apiuser.Username, apiuser.Email, apiuser.Key)
			if err != nil {
				c.JSON(400, errors.New(errorCodes.TOKENERROR).Error())
			}
			queriedShop, err := database.QueryShopByApiKey(models.SHOPLIST, apiuser.Key)
			if err != nil {
				shop.ApiKey = apiuser.Key
				shop.Name = queriedShop.Name
				models.SHOPLIST = append(models.SHOPLIST, *shop)
				c.JSON(200, token)
			} else {
				log.Print(err)
				c.JSON(400, errorCodes.USERALREADYLOGGEDIN)
				c.JSON(400, err)
			}
		}
	}
}
func MainPage(c *gin.Context) {
	apiUSER := new(models.APIUSER)
	err := c.ShouldBind(&apiUSER)
	if err != nil {
		c.JSON(400, errorCodes.COULDNOTBIND)
	} else {
		username, email, apikey, err := webtoken.GetValidTokenValue(apiUSER.Token)
		switch {
		case err == nil:

		case err.Error() == errorCodes.TOKENERROR:
			_, _, apikey := webtoken.GetInvalidTokenValue(apiUSER.Token)
			if apikey == "" {
				log.Fatal(err)
			}
			removed := database.RemoveShop(apikey)
			if !removed {
				log.Print("error while deleting")
			}
			c.JSON(400, errorCodes.TOKENERROR)
		case err.Error() == errorCodes.TOKENEXPIRED:
			log.Print(errorCodes.TOKENEXPIRED)
			_, _, apikey := webtoken.GetInvalidTokenValue(apiUSER.Token)
			if apikey == "" {
				log.Fatal(err)
			}
			removed := database.RemoveShop(apikey)
			if !removed {
				log.Print("error while deleting")
			}
			c.JSON(403, errorCodes.TOKENEXPIRED)
		}

	}
}
func RenewApiKey(c *gin.Context) {
	var (
		apiUser = new(models.APIUSER)
	)
	err := c.ShouldBind(&apiUser)
	if err != nil {
		c.JSON(400, errorCodes.COULDNOTBIND)
	} else {
		newtoken, err := webtoken.CreateToken(webtoken.GetInvalidTokenValue(apiUser.Token))
		if err != nil {
			c.JSON(400, errorCodes.TOKENERROR)
		} else {
			c.JSON(200, newtoken)
		}

	}
}
