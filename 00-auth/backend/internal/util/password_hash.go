package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)




func HashPassword (password string) (string, error) {
	hashedPassword, err := 	bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
 if err != nil {
	 return "",fmt.Errorf("error happen during hashing :%w",err) 
 }
 return string(hashedPassword), nil
}


func CheckPassword(password,hashedPassword string) bool {

 if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err !=nil {
	 return false 
 }
 return true
}
