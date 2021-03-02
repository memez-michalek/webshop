package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"tback/controllers"
	"tback/helper"

	"github.com/gin-gonic/gin"
)

func initialConnection() {
	marshalled, err := json.Marshal(&helper.SHOPAPISECRET)
	if err != nil {
		log.Print("could not marshal api login credentials")
	}
	resp, err := http.Post("http://web:8000/api/login", "application/json", bytes.NewReader(marshalled))
	if err != nil {
		log.Print(err)
		log.Fatal("initial request error")
	}
	if resp.StatusCode != 200 {
		log.Fatal("could not login to api")

	}
	if resp.StatusCode == 403 {
		marshalled, err := json.Marshal(helper.SHOPAPIKEY)
		if err != nil {
			log.Print(err)
			log.Print("could not marshall")
		}
		resp, err := http.Post("http://webshop:8000/api/renew", "application/json", bytes.NewReader(marshalled))
		if err != nil {
			log.Print(err)
			log.Print("error occurred")
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
			log.Print("could not read body")
		}
		log.Print(string(data))
		helper.SHOPAPIWEBTOKEN = string(data)

	}

	data, err := ioutil.ReadAll(resp.Body)
	helper.SHOPAPIWEBTOKEN = string(string(data))

}

func main() {
	initialConnection()
	router := gin.Default()
	router.POST("/init", controllers.MainSiteController)
	router.POST("/login", controllers.LoginController)
	router.POST("/logout", controllers.LoginController)
	router.POST("/register", controllers.RegisterController)
	router.Run(":8080")
}
