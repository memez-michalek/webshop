package database

import (
	"api/helpers"
	"api/models"
	"context"
	"errors"
	"log"

	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func initDB(db string, collection string) *mongo.Collection {

	options := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), options)
	helpers.ErrorChecker(err)

	err = client.Ping(context.TODO(), nil)
	helpers.ErrorChecker(err)

	output := client.Database(db).Collection(collection)
	return output

}

func InsertUser(email string, username string, password string) bool {
	var (
		collection = initDB("user", "credentials")
		user       = new(models.User)
	)
	filter := bson.D{{"name", username}, {"email", email}}
	e := collection.FindOne(context.TODO(), filter).Decode(&user)
	if e != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		helpers.ErrorChecker(err)
		id := ksuid.New().String()
		log.Print(id)
		user.Id = id
		user.Username = username
		user.Email = email
		user.Password = hash
		user.ApiKeys = []string{}

		_, err = collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Print("not inserted")
			helpers.ErrorChecker(err)
			return false
		}
		log.Print("inserted")
		return true

	} else {
		log.Print("user already exists")
		return false
	}

}

func LoginUser(email string, username string, password string) bool {
	var (
		collection = initDB("user", "credentials")
		user       = new(models.User)
		filter     = bson.D{{"username", username}, {"email", email}}
	)
	log.Print(filter)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print("error user does not exist")
		return false
	} else {
		hashedpw := user.Password
		err = bcrypt.CompareHashAndPassword(hashedpw, []byte(password))
		if err != nil {
			log.Print(err)
			return false
		} else {
			log.Print("logged in")
			return true
		}
	}
}
func AddApi(email string, username string) bool {
	var (
		collection = initDB("user", "credentials")
		user       = new(models.User)
		filter = bson.D{{"username", username}, {"email", email}}
	)
	log.Print(filter)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print("user does not exist")
		return false

	} else {
		id := ksuid.New().String()
		user.ApiKeys = append(user.ApiKeys, id)
		update := bson.M{"$set": bson.M{"apikeys": user.ApiKeys}}
		e := collection.FindOneAndUpdate(context.TODO(), filter, update).Err()
		if e != nil {
			log.Print("db has not been updated")

			return false
		} else {
			log.Print("updated")
			return true
		}

	}
}
func GetApiKeys(email string, username string) ([]string, error) {
	var(
	collection = initDB("user","credentials")
	user = new(models.User)
	filter = bson.M{"username": username, "email": email}
	)
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		log.Print("could not find user")
		return []string{}, errors.New("could not find any user with such credentials")
	} else {
		log.Print("users apis found")
		return user.ApiKeys, nil
	}
}
