package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	var password string
	fmt.Println("Your Password should not contain any spaces")
	fmt.Print("Please enter your password: ")
	fmt.Scanln(&password)
	hash, _ := HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hashed Password:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)

	fmt.Println("decrypt", hash)

}

// Decryption of hashed password algorithm is comming soon....
