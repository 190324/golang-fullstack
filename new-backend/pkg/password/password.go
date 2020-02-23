package password

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) (string, err) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash), err
}

func ComparePassword(p string, hash string) bool {
	err = bcrypt.CompareHashAndPassword([]byte(hasg), []byte(p))
	if err != nil {
		return false
	} else {
		return true
	}
}
