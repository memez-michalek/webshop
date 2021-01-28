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
	Category string `json:"category"`
	Name     string `json:"name"`
}
type QUERYShop struct {
	Name string
	ID   string
}

type SHOP struct {
	ID    string    `json:"shop_id"`
	Name  string    `json:"name"`
	ITEMS []Product `json:"ITEMS"`
}

type APIUSER struct {
	Token string `json:"key"`
}
