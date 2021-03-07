package controllers

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"

	"net/http"
	"tback/handlers"
	"tback/helper"
	"tback/model"

	"github.com/gin-gonic/gin"
)

func MainSiteController(c *gin.Context) {
	var (
		data = `{"key": ` + helper.SHOPAPIWEBTOKEN + `}`
	)
	marshalled := []byte(data)
	log.Print(string(marshalled))
	resp, err := http.Post("http://web:8000/api/main", "application/json", bytes.NewReader(marshalled))
	log.Print(err)
	switch {
	case err != nil:
		log.Print("api request error")
		c.JSON(404, "error when sending")
	case resp.StatusCode != 200:
		buff := []byte{}
		_, err := io.ReadFull(resp.Body, buff)
		if err != nil {
			log.Print(err)
		}
		log.Print(resp.StatusCode)
		log.Print(string(buff))
		log.Print("could not get site products")
		c.JSON(404, "could not get products")
	case resp.StatusCode == 200:
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
			log.Print("could not read body")
		}
		log.Print(string(data))
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.JSON(200, string(data))
	}
}
func LoginController(c *gin.Context) {
	var (
		user = new(model.User)
	)
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	err := c.ShouldBind(&user)
	if err != nil {
		log.Print("access denied")
		c.JSON(403, err)
	} else {
		err := handlers.CheckLoginData(*user)
		if err != nil {
			log.Print("login data wasnt verified: ", err)
			c.JSON(404, err)
		} else {

			token, err := helper.CreateToken(user.Email, user.Password)
			if err != nil {
				log.Print("could not create login token", err)
				c.JSON(403, err)
			} else {
				model.VALIDJWTTOKENS[user.Email] = token
				c.JSON(200, "logged in")
			}
		}
	}

}
func RegisterController(c *gin.Context) {
	var (
		user = new(model.User)
	)
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	log.Print(user)
	err := c.ShouldBind(&user)
	if err != nil {
		log.Print("could not bind to model")
		c.JSON(400, "could not bind")
	} else {

		err := handlers.RegisterUser(*user)
		if err != nil {
			log.Print("bad user credentials", err)
			c.JSON(403, "could not log in")
		}
		token, err := helper.CreateToken(user.Email, user.Password)
		if err != nil {
			log.Print("could not create token")
		} else {
			model.VALIDJWTTOKENS[user.Email] = token
			c.JSON(200, "user registered")
		}

	}

}
func LogoutController(c *gin.Context) {
	var (
		user = new(model.VerifiedUser)
	)
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	email, _, err := helper.GetTokenValue(user.Webtoken)
	if err != nil {
		log.Print("could not get the value of token")
		c.JSON(401, "could not get value of token")
	}
	delete(model.VALIDJWTTOKENS, email)
	c.JSON(200, "logged out ")
}
