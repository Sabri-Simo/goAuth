package service

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func HashPassword(passwrod string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(passwrod), bcrypt.DefaultCost)
	return string(hashPassword), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsvalidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
