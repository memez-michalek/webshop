package main

import (
	"api/controllers"
	"api/views"
	"context"
	

	"github.com/gin-gonic/gin"
)

var (
	ctx = context.Background()
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	website := router.Group("/")
	{

		website.GET("/", views.IndexView)
		website.GET("/login", views.LoginView)
		website.POST("/login", controllers.LoginHandler)
		website.GET("/register", views.RegisterView)
		website.POST("/register", controllers.RegisterHandler)
		website.GET("/logout", controllers.LogoutHandler)
		website.GET("/getapi", controllers.GetApiHandler)
	}

	router.Run(":8000")
}
