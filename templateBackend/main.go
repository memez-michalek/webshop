package templateBackend

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default
	router.POST("/login")
	router.POST("/logout")
	router.POST("/register")
	router.Run(":8080")
}
