package views

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexView(c *gin.Context){
	username, err :=c.Cookie("username")
	if err != nil{
		log.Print(err, username)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Username" : "jebac disa",
	})

}




