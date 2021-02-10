package database

import (
	"api/errorCodes"
	"api/models"
	"api/webtoken"
	"context"
	"errors"
	"log"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
)

func QueryShop(queryshop models.QUERYShop) (models.SHOP, error) {
	var (
		shop       = new(models.SHOP)
		collection = initDB("DATABASE", "SHOPS")
		filter     = bson.M{"name": queryshop.Name, "shop_id": queryshop.ID}
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&shop)
	if err != nil {
		log.Print(err)
		return *shop, errors.New(errorCodes.COULDNOTFINDSHOP)
	}
	return *shop, nil
}
func GetOrder(shop models.SHOP, orderid string) (models.Order, error) {
	for _, v := range shop.Orders {
		if v.Id == orderid {
			return v, nil
		}
	}
	return models.Order{}, errors.New(errorCodes.ORDERDOESNOTEXIST)
}
func RemoveOrder(order models.Order, shop models.SHOP) ([]models.Order, error) {
	for i, v := range shop.Orders {
		if v.Id == order.Id {
			shop.Orders = append(shop.Orders[:i], shop.Orders[i+1:]...)
			return shop.Orders, nil
		}
	}
	return []models.Order{}, errors.New(errorCodes.ORDERDOESNOTEXIST)
}

func QueryOrder(orderfilter models.OrderFilter) (models.Order, error) {

	_, _, apikey, err := webtoken.GetValidTokenValue(orderfilter.Webtoken)
	if err != nil {
		log.Print(err)
		return models.Order{}, err
	}
	queryshop, err := QueryShopByApiKey(models.SHOPLIST, apikey)
	if err != nil {
		log.Print(err)
		return models.Order{}, err
	}
	shop, err := QueryShop(queryshop)
	if err != nil {
		log.Print(err)
		return models.Order{}, err
	}
	order, err := GetOrder(shop, orderfilter.OrderId)
	if err != nil {
		log.Print(err)
		return models.Order{}, err
	}
	return order, nil
}
func AddOrder(shopId string, order models.Order) error {
	var (
		filter     = bson.M{"shop_id": shopId}
		collection = initDB("DATABASE", "SHOPS")
		shop       = new(models.SHOP)
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&shop)
	if err != nil {
		log.Print(err)
		return err
	} else {
		shop.Orders = append(shop.Orders, order)
		update := bson.M{"$set": bson.M{"Orders": shop.Orders}}
		err = collection.FindOneAndUpdate(context.TODO(), filter, update).Err()
		if err != nil {
			log.Print(err)
			return err
		}
		return nil
	}
}

func MakeOrder(queryOrder *models.QueryOrder) (models.Order, error) {
	var (
		order = new(models.Order)
	)
	queryShop, err := GetQueryShop(*queryOrder)

	shop, err := QueryShop(queryShop)
	if err != nil {
		log.Print(err)
		return *order, err
	}

	products, err := QueryProductsByProductIds(shop.ITEMS, queryOrder.ProductIds)
	if err != nil {
		log.Print(err)
		return *order, err
	}
	order.Id = ksuid.New().String()
	order.Credentials = queryOrder.Credentials
	order.Products = products
	return *order, nil
}
func GetQueryShop(queryOrder models.QueryOrder) (models.QUERYShop, error) {

	_, _, apikey, err := webtoken.GetValidTokenValue(queryOrder.Webtoken)
	if err != nil {
		log.Print(errorCodes.TOKENEXPIRED)
		return models.QUERYShop{}, err
	}
	queryshop, err := QueryShopByApiKey(models.SHOPLIST, apikey)
	if err != nil {
		log.Print(err)
		return models.QUERYShop{}, err
	}
	return queryshop, nil
}
func DeleteOrder(order models.Order, Webtoken string) error {
	var (
		collection = initDB("DATABASE", "SHOPS")
		shop       = new(models.SHOP)
	)
	_, _, apikey, err := webtoken.GetValidTokenValue(Webtoken)
	if err != nil {
		log.Print(errorCodes.TOKENEXPIRED)
		return err
	}
	queryShop, err := QueryShopByApiKey(models.SHOPLIST, apikey)
	if err != nil {
		log.Print(errorCodes.TOKENEXPIRED)
		return err
	}
	err = collection.FindOne(context.TODO(), bson.M{"shop_id": queryShop.ID}).Decode(&shop)
	if err != nil {
		log.Print(err)
		return err
	} else {
		updatedOrders, err := RemoveOrder(order, *shop)
		if err != nil {
			log.Print(err)
			return err
		}
		shop.Orders = updatedOrders
		err = collection.FindOneAndUpdate(context.TODO(), bson.M{"shop_id": queryShop.ID}, bson.M{"$set": bson.M{"Orders": shop.Orders}}).Err()
		if err != nil {
			log.Print(err)
			return err
		}
		return nil
	}

}
