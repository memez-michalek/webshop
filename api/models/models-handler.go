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
	jwt.StandardClaims
}
