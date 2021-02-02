package models

import "github.com/dgrijalva/jwt-go"

var SHOPLIST = []QUERYShop{}

type User struct {
	Id       string
	Username string
	Email    string
	Password []byte
	ApiKeys  map[string]string
}

type JWTAUTH struct {
	Username string
	Email    string
	Apikey   string
	jwt.StandardClaims
}
type APILOGIN struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Key      string `json:"key"`
}

type Product struct {
	ID          string `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
}
type QUERYShop struct {
	Name string
	ID   string
}

type SHOP struct {
	SHOP_ID string    `json:"shop_id"`
	Name    string    `json:"name"`
	ITEMS   []Product `json:"ITEMS"`
}

type APIUSER struct {
	Token string `json:"key"`
}
type APIUSERADDPRODUCTS struct {
	Token string    `json:"key"`
	ITEMS []Product `json:"ITEMS"`
}
type QueryProduct struct {
	Token     string `json:"key"`
	ProductId string `json:"product_id"`
}
