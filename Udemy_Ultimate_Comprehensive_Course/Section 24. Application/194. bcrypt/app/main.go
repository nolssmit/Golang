package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unencrypted password: ", s)
	fmt.Println("Encrypted password: ", string(bs))
    
	// Test if logi password is valid
	loginpassword := `password123`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpassword))
	if err != nil {
		fmt.Println("YOU CAN't LOGIN")
		return
	}
	fmt.Println("Login successful")

}

// Intall package: go get golang.org/x/crypto/bcrypt
