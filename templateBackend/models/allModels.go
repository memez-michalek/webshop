package allModels

import "github.com/dgrijalva/jwt-go"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type Auth struct {
	Email    string
	Password string
	Claims   jwt.StandardClaims
}
