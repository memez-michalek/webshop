package database

import (
	"api/models"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func ApiLogin(email string, username string, key string) error {
	var(
	collection = initDB("user","credentials")
	user = new(models.User)
	filter = bson.M{"username": username, "email": email}
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print("user does not exist")
		return errors.New("user does not exist")

	} else {
		for _, v := range user.ApiKeys {
			if v == key {
				return nil
			}
		}
		return errors.New("could not find such api key")
	}
}
func GetMainSiteProducts(){
	collection := initDB("")

}





