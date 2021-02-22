package templateBackend

import (
	"api/helpers"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

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
	router.POST("/init")
	router.POST("/login")
	router.POST("/logout")
	router.POST("/register")
	router.Run(":8080")
}
