package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       string
	Username string
	Email    string
	Password []byte
	ApiKeys  []string
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
