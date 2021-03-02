package controllers

import (
	"api/database"
	"api/errorCodes"
	"api/models"
	"api/webtoken"
	"errors"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
)

func LogInToApi(c *gin.Context) {
	var (
		apiuser = new(models.APILOGIN)
		shop    = new(models.QUERYShop)
	)
	err := c.ShouldBindJSON(&apiuser)
	log.Print(err)

	if err != nil {
		c.JSON(400, errorCodes.COULDNOTBIND)

	} else {
		shopname, err := database.ApiLogin(apiuser.Email, apiuser.Username, apiuser.Key)
		if err != nil {
			c.JSON(400, "wrong credentials: "+err.Error())
		} else {
			token, err := webtoken.CreateToken(apiuser.Username, apiuser.Email, apiuser.Key)
			if err != nil {
				c.JSON(400, errors.New(errorCodes.TOKENERROR).Error())
			}
			_, err = database.QueryShopByApiKey(models.SHOPLIST, apiuser.Key)
			if err != nil {
				shop.ID = apiuser.Key
				shop.Name = shopname
				models.SHOPLIST = append(models.SHOPLIST, *shop)
				log.Print("shoplist", models.SHOPLIST)
				log.Print(token)
				c.JSON(200, token)
			} else {
				log.Print(err)
				c.JSON(403, token)
			}
		}
	}
}
func MainPage(c *gin.Context) {
	var (
		apiUSER = new(models.APIUSER)
	)
	err := c.ShouldBind(&apiUSER)
	if err != nil {
		log.Print(apiUSER)
		c.JSON(400, errorCodes.COULDNOTBIND)
	} else {

		log.Print("non decoded token", apiUSER.Token)
		typ := reflect.TypeOf(apiUSER.Token)
		log.Print(typ)

		_, _, apikey, err := webtoken.GetValidTokenValue(string(apiUSER.Token))
		log.Print(err, "<= error     apikey =>", apikey)
		switch {
		case err == nil:

			log.Print(models.SHOPLIST)
			shop, err := database.QueryShopByApiKey(models.SHOPLIST, apikey)
			if err != nil {
				c.JSON(400, errorCodes.SHOPDOESNOTEXIST)
			}
			products, err := database.GetMainSiteProducts(shop)
			if err != nil {
				log.Print(err)
				c.JSON(400, errorCodes.SHOPDOESNOTEXIST)
			} else {

				c.JSON(200, products)
			}
		case err.Error() == errorCodes.TOKENERROR:
			_, _, apikey := webtoken.GetInvalidTokenValue(apiUSER.Token)
			if apikey == "" {
				c.JSON(404, "Fatal error ")
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
				c.JSON(404, "Fatal error ")
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
func LogOutFromAPI(c *gin.Context) {
	var (
		apiWebtoken = new(models.APIUSER)
	)
	err := c.ShouldBind(&apiWebtoken)
	if err != nil {
		log.Print(errorCodes.COULDNOTBIND)
		c.JSON(400, err)
	} else {
		_, _, apikey, err := webtoken.GetValidTokenValue(apiWebtoken.Token)
		if err != nil {
			log.Print(errorCodes.TOKENEXPIRED)
			c.JSON(400, err)
		} else {
			removed := database.RemoveShop(apikey)
			if !removed {
				log.Print(errorCodes.SHOPDOESNOTEXIST)
				c.JSON(400, errorCodes.SHOPDOESNOTEXIST)
			}
		}
	}

}

func CreateShopController(c *gin.Context) {
	var (
		user = new(models.APIUSER)
	)
	err := c.ShouldBind(&user)
	if err != nil {
		log.Print(err)
		c.JSON(400, errorCodes.COULDNOTBIND)
	} else {
		_, _, apikey, err := webtoken.GetValidTokenValue(user.Token)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		}
		shop, err := database.QueryShopByApiKey(models.SHOPLIST, apikey)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		}
		err = database.CreateNewShop(shop)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		} else {
			c.JSON(200, "Shop created")
		}

	}
}
func InsertProductsIntoShopController(c *gin.Context) {
	var (
		user = new(models.APIUSERADDPRODUCTS)
	)
	err := c.ShouldBind(&user)
	log.Print("malik montana", user)
	if err != nil {
		log.Print(errorCodes.COULDNOTBIND)
		c.JSON(400, err)
	} else {

		_, _, apikey, err := webtoken.GetValidTokenValue(user.Token)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		}
		shop, err := database.QueryShopByApiKey(models.SHOPLIST, apikey)
		if err != nil {
			c.JSON(400, err)
		}
		err = database.InsertSiteProducts(shop, user.ITEMS)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		} else {

			c.JSON(200, "inserted")
		}
	}
}
func GetItemDetails(c *gin.Context) {
	var (
		product = new(models.QueryProduct)
	)

	err := c.ShouldBind(&product)
	if err != nil {
		log.Print(errorCodes.COULDNOTBIND)
		c.JSON(400, errorCodes.COULDNOTBIND)
	} else {

		_, _, apiKey, err := webtoken.GetValidTokenValue(product.Token)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		}
		shop, err := database.QueryShopByApiKey(models.SHOPLIST, apiKey)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		}
		productDetails, err := database.GetItemDetails(shop, product.ProductId)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		} else {
			c.JSON(200, productDetails)
		}
	}

}
