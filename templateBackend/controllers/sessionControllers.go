package controllers

import ("github.com/gin-gonic/gin"

"templateBackend/models"
 "net/http"
"templateBackend/handlers")

func MainSiteController(c *gin.Context){
	marshalled, err := json.Marshal(helpers.SHOPAPIWEBTOKEN)
	resp, err := http.Post("/api/main", bytes.NewReader(marshalled))
	switch {
	case err != nil:
		log.Print("api request error")
		c.JSON(404, "error when sending")
	case resp.Code != 200:
		log.Print("could not get site products")
		c.JSON(404, "could not get products")
	case resp.Code == 200:
		c.JSON(200, resp.Body)
	}
}
func LoginController(c *gin.Context){
	var(
		login := new(models.User)
	)
	err := c.ShouldBind(&login)
	if err != nil {
		log.Print("access denied")
		c.JSON(403, err)
	}else{
		err := login.CheckLoginData()
		if err != nil {
			log.Print("login data wasnt verified: ", err)
			c.JSON(404, err)
		}else{
			_ , exists := models.VALIDJWTTOKENS[login.Email]
			if exists{
				log.Print("you are already logged in")
				c.JSON(403, "already logged in")
			}
			token, err  := helpers.CreateToken(login.Email, login.Password)	
			if err != nil{
				log.Print("could not create login token", err)
				c.JSON(403, err)
			}
			models.VALIDJWTTOKENS[login.Email] = token
			c.JSON(200, "logged in")

		}
	}

}
func RegisterController(c *gin.Context){
	var(
		user := new(models.User)
	)
	err := c.ShouldBind(&user)
	if err != nil{
		log.Print("could not bind to model")
		c.JSON(400, "could not bind")
	}else{
		_ , exists := models.VALIDJWTTOKENS[user.Email]
			if exists{
				log.Print("you are already logged in")
				c.JSON(403, "already logged in")
			}
	
		err := user.RegisterUser()
		if err != nil{
			log.Print("bad user credentials", err)
			c.JSON(403, "could not log in")
		}
		token, err := helpers.CreateToken(user.Email, user.Password)
		if err != nil{
			log.Print("could not create token")
		}
		models.VALIDJWTTOKENS[user.Email] = token
		c.JSON(200, "user registered")

	}	

}
