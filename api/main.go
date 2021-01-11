package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/memez-Michalek/"
)

var (
	ctx = context.Background()
)

func main() {
	router := gin.Default()
	router.Get("/", views.IndexView)

	router.run(":8080")

}
