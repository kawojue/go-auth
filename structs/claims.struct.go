package structs

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
