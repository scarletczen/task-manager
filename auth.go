package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func WithJWTAuth(handlerFunc http.HandlerFunc, store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetTokenFromRequest(r)
		token, err := validateJWT(tokenString)
		if err != nil {
			log.Println("failed to authenticate token")
			permissionDenied(w)
			return
		}
		if !token.Valid {
			log.Println("failed to authenticate token")
			permissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		_, err = store.GetUserById(userId)
		if err != nil {
			log.Println("failed to get user")
			permissionDenied(w)
			return
		}
	}
}


func validateJWT(t string) (*jwt.Token, error) {
	secret := Envs.JWTSecret
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
		Error: fmt.Errorf("permission denied").Error(),
	})
}

func HashPassword(pw string) (string, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CreateJWT(secret []byte, userId int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"userId": strconv.Itoa(int(userId)),
		"expiresAt": time.Now().Add(time.Hour * 24 *120).Unix(),
	})	
	tokenString, err := token.SignedString(secret)
	if err!=nil{
		return "", err
	}
	return tokenString, nil
}