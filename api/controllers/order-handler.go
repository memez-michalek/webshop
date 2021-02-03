package controllers

import (
	"api/database"
	"api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func MakeOrder(c *gin.Context) {
	var (
		queryOrder = new(models.QueryOrder)
	)
	err := c.ShouldBind(&queryOrder)
	if err != nil {
		log.Print(err)
		c.JSON(400, err)
	} else {

		database.MakeOrder(queryOrder)
		order, err := database.MakeOrder(queryOrder)
		if err != nil {
			log.Print(err)
			c.JSON(400, err)
		} else {
			queryshop, err := database.GetQueryShop(*queryOrder)
			if err != nil {
				log.Print(err)
				c.JSON(400, err)
			}
			err = database.AddOrder(queryshop.ID, order)
			if err != nil {
				log.Print(err)
				c.JSON(400, err)
			} else {
				c.JSON(200, "inserted")
			}
		}
	}
}
