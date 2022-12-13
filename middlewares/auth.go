package middlewares

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) string {
	passwordByte, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newPassword := string(passwordByte)
	return newPassword
}