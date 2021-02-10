package database

import (
	"api/errorCodes"
	"api/models"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func QueryShopByApiKey(shops []models.QUERYShop, apiKey string) (models.QUERYShop, error) {
	for _, v := range shops {
		log.Print(v)
		if v.ID == apiKey {
			return v, nil
		}
	}
	return models.QUERYShop{}, errors.New(errorCodes.SHOPDOESNOTEXIST)
}
func QueryProductsByProductIds(products []models.Product, ID []string) ([]models.Product, error) {
	var (
		queriedproducts = make([]models.Product, 0)
	)
	log.Print(products, ID)
	for _, id := range ID {
		for _, v := range products {
			log.Print(id)
			if v.ID == id {
				queriedproducts = append(queriedproducts, v)
			}
		}
	}
	if len(queriedproducts) > 0 {
		return queriedproducts, nil
	}
	return queriedproducts, errors.New(errorCodes.COULDNOTFINDPRODUCTS)
}

func ApiLogin(email string, username string, key string) (string, error) {
	var (
		collection = initDB("user", "credentials")
		user       = new(models.User)
		filter     = bson.M{"username": username, "email": email}
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print(errorCodes.USERDOESNOTEXIST)
		return "", errors.New(errorCodes.USERDOESNOTEXIST)

	} else {
		for i, v := range user.ApiKeys {
			if i == key {
				return v, nil
			}
		}
		return "", errors.New(errorCodes.COULDNOTFINDAPIKEY)
	}
}
func RemoveShop(apiKey string) bool {
	for i, v := range models.SHOPLIST {
		if v.ID == apiKey {
			models.SHOPLIST = append(models.SHOPLIST[:i], models.SHOPLIST[i+1:]...)
			return true
		}
	}
	return false
}

//NOTE OUTPUT TYPE MIGHT / WILL CHANGE
func GetMainSiteProducts(queryShop models.QUERYShop) ([]models.Product, error) {
	var (
		collection = initDB("DATABASE", "SHOPS")
		filter     = bson.M{"shop_id": queryShop.ID, "name": queryShop.Name}
		shop       = new(models.SHOP)
		product    = new(models.Product)
	)
	shop.ITEMS = append(shop.ITEMS, *product)
	err := collection.FindOne(context.TODO(), filter).Decode(&shop)
	if err != nil {
		log.Print(err)
		log.Print(errorCodes.COULDNOTBIND)
		return shop.ITEMS, errors.New(errorCodes.COULDNOTBIND)
	} else {
		return shop.ITEMS, nil
	}
}

func InsertSiteProducts(queryShop models.QUERYShop, products []models.Product) error {
	var (
		collection = initDB("DATABASE", "SHOPS")
		shop       = new(models.SHOP)
		filter     = bson.M{"name": queryShop.Name, "shop_id": queryShop.ID}
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&shop)
	if err != nil {

		return errors.New(errorCodes.SHOPDOESNOTEXIST)
	} else {
		shop.ITEMS = append(shop.ITEMS, products...)

		update := bson.M{"$set": bson.M{"ITEMS": shop.ITEMS}}

		err := collection.FindOneAndUpdate(context.TODO(), filter, update).Err()
		if err != nil {
			log.Print(err)
			return errors.New(errorCodes.COULDNOTINSERTINTODB)
		}

		return nil
	}
}
func CreateNewShop(QueryShop models.QUERYShop) error {
	var (
		collection = initDB("DATABASE", "SHOPS")
		shop       = new(models.SHOP)
		//items      = make([]models.Product, 0)
		filter = bson.M{"shop_id": QueryShop.ID, "name": QueryShop.Name}
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&shop)
	if err != nil {
		shop.SHOP_ID = QueryShop.ID
		shop.Name = QueryShop.Name
		_, err := collection.InsertOne(context.TODO(), shop)
		if err != nil {
			log.Print(err)
			return errors.New(errorCodes.SHOPWASTNINSERTED)
		}
		return nil
	}
	return errors.New(errorCodes.SHOPALREADYEXISTS)
}
func GetItemDetails(QueryShop models.QUERYShop, productId string) (models.Product, error) {
	var (
		collection = initDB("DATABASE", "SHOPS")
		filter     = bson.M{"shop_id": QueryShop.ID, "name": QueryShop.Name}
		shop       = new(models.SHOP)
		product    = new(models.Product)
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&shop)
	if err != nil {
		log.Print(errorCodes.COULDNOTBIND)
		log.Print(err)
		return *product, err
	} else {
		foundproduct, err := QueryProductsByProductIds(shop.ITEMS, []string{productId})
		if err != nil {
			log.Print(err)
			return *product, err
		}
		return foundproduct[0], nil

	}
}
