package repository

import (
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

func checkLength(name string) bool {
	return len(name) > 1
}

// Hash Passowrd Functionality
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func checkHashPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
