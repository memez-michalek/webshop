package database

import (
	"api/models"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func QueryShopByApiKey(shops []models.Shop, apiKey string) (models.Shop, error) {
	for _, v := range shops {
		log.Print(v)
		if v.ApiKey == apiKey {
			return v, nil
		}
	}
	return models.Shop{}, errors.New("shop does not exist")

}
func ApiLogin(email string, username string, key string) error {
	var (
		collection = initDB("user", "credentials")
		user       = new(models.User)
		filter     = bson.M{"username": username, "email": email}
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print("user does not exist")
		return errors.New("user does not exist")

	} else {
		for i := range user.ApiKeys {
			if i == key {
				return nil
			}
		}
		return errors.New("could not find such api key")
	}
}

/*
NOTE OUTPUT TYPE MIGHT / WILL CHANGE
func (shop models.Shop) GetMainSiteProducts() (query []string, err error) {

}
*/
