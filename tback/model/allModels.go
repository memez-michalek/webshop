package model

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type AdminUser struct {
	Email    string `json:"email"`
	Secret   string `json:"secret"`
	Password string `json:"password"`
}

type Auth struct {
	Email    string
	Password string
	jwt.StandardClaims
}
type VerifiedUser struct {
	Webtoken string
}

type Order struct {
	ProductIds  []string    `json:"productids"`
	Credentials Credentials `json:"credentials"`
}
type ApiOrder struct {
	Webtoken    string      `json:"webtoken"`
	ProductIds  []string    `json:"productids"`
	Credentials Credentials `json:"credentials"`
}

type Credentials struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}
type GetOrderByEmail struct {
	Email string `json:"email"`
}

type OrderFilter struct {
	Webtoken string   `json:"webtoken"`
	OrderId  []string `json:"orderid"`
}
type Filter struct {
	OrderId []string `json:"orderid"`
}

type Ord struct {
	Id          string      `json:"ID"`
	Products    []Product   `json:"products"`
	Credentials Credentials `json:"credentials"`
}

type Product struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Image       string `json:"imageUrl"`
}

func (apiOrder *ApiOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Webtoken    string      `json:"webtoken"`
		ProductIds  []string    `json:"productids"`
		Credentials Credentials `json:"credentials"`
	}{
		Webtoken:    apiOrder.Webtoken,
		ProductIds:  apiOrder.ProductIds,
		Credentials: apiOrder.Credentials,
	})
}
