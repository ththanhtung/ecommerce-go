package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func VerifyPassword(hashedPassword, userPassword string) (bool, error){
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword)); err != nil {
		return false, err
	}

	return true, nil
}