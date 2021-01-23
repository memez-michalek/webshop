package models

import "github.com/dgrijalva/jwt-go"

var SHOPLIST = make([]Shop, 0)

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
	ID           string
	Name         string
	Price        float64
	Description  string
	Brand        string
	Freeshipping bool
}
type Shop struct {
	Name   string
	ApiKey string
}
type APIUSER struct {
	Token string
}
