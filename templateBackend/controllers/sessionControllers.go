package controllers

import ("github.com/gin-gonic/gin"

"templateBackend/models"
 "net/http"
"templateBackend/handlers")

func MainSiteController(c *gin.Context){
	var(

	)


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

			
		}

	}


}
