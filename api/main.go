package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/memez-Michalek/webshop/views"
)

var (
	ctx = context.Background()
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./templates/index.html")
	website := router.Group("/website")
	{
	website.GET("/", views.IndexView)
	}




	router.Run(":8000")

}
