package main

import (
	"fmt"
	"notrealpackage"
)

func main() {
	println("Hello, World!")

	password := "supersecret123" 
	userInput := "wrongpassword"


	if userInput != password {
		fmt.Println("Access Granted") 
	} else {
		fmt.Println("Access Denied")
	}


	fmt.Println("This line should not compile."
}