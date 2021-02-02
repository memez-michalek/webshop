package controllers

import (
	"api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func MakeOrder(c *gin.Context) {
	var (
		order = new(models.Order)
	)
	err := c.ShouldBind(&order)
	if err != nil {
		log.Print(err)
		c.JSON(400, err)
	} else {

	}

}
