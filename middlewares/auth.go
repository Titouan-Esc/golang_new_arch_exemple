package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtKey = []byte("user")

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func HasPassword(password string) string {
	passwordByte, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newPassword := string(passwordByte)
	return newPassword
}

func ValidateEncrypt(password, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	return true
}

func GenerateJWT(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(240 * time.Hour)
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)

	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return fmt.Errorf("Couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return fmt.Errorf("Token expired")
	}
	return
}
