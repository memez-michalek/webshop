package templateBackend

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"tback/controllers"
	"tback/helpers"

	"github.com/gin-gonic/gin"
)

func main() {
	helpers.SHOPAPISECRET
	marshalled, err := json.Marshal(&helpers.SHOPAPISECRET)
	if err != nil {
		log.Print("could not marshal api login credentials")
	}
	resp, err := http.Post("http://localhost:8000/api/login", bytes.NewReader(marshalled))
	if err != nil {
		log.Print("initial request error")
	}
	if resp.Code != 200 {
		log.Fatal("could not login to api")
	}
	helpers.SHOPAPIWEBTOKEN = resp.Body

	router := gin.Default()
	router.POST("/init", controllers.MainSiteController)
	router.POST("/login", controllers.LoginHandler)
	router.POST("/logout", controllers.LogoutHandler)
	router.POST("/register", controllers.RegisterHandler)
	router.Run(":8080")
}
