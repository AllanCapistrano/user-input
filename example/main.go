package main

import (
	"fmt"

	"github.com/allancapistrano/user-input"
)

func main() {
	input := userinput.GetUserInput("Enter the text to be displayed: ")

	fmt.Println("\nInput: " + input)
}
