package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Email    string `json:"email"`
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
