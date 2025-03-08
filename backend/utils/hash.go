package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	salt := bcrypt.DefaultCost
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return "", err
	}
	
	return string(hashed), nil
}