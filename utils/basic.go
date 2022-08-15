package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return string(result)
}

func ValidatePassword(hashPassword *string, password string) error {
	hash := *hashPassword
	fmt.Println(hash)
	fmt.Println(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
