package JWT

import (
	"time"
	"errors"
	"api/config"
	"github.com/dgrijalva/jwt-go"
)
type CustomClaims struct {
	ID     uint   `json:"id"`
 	Name   string `json:"name"`
 	Email  string `json:"email"`
 	jwt.StandardClaims
}

func GenerateToken(id uint, name, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims(id, name, email))
	tokenString, err := token.SignedString([]byte(config.Config("jwt.secret").(string))) 
	return tokenString, err
}

func Check(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method: " + token.Header["alg"].(string))
		}
		return []byte(config.Config("jwt.secret").(string)), nil
	})
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
		return true
	} else {
		return false
	}
}

func ParseToken(tokenStr string) *CustomClaims {
	token, _ := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config("jwt.secret").(string)), nil
	})
	customClaims, _ := token.Claims.(*CustomClaims)
	return customClaims
}


func setClaims(id uint, name, email string) CustomClaims {
	return CustomClaims{
		id,
		name,
		email,
		jwt.StandardClaims{
			Audience: config.Config("jwt.aud").(string),
			ExpiresAt: time.Now().Unix() + int64(config.Config("jwt.exp").(int)),
			IssuedAt: time.Now().Unix(),
			Issuer: config.Config("jwt.iss").(string),
			NotBefore: time.Now().Unix(),
			Subject: config.Config("jwt.sub").(string),
		},
	} 
}