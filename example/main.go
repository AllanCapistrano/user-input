package main

import (
	"fmt"

	"github.com/allancapistrano/user-input"
)

func main() {
	input := userinput.GetUserInput("Enter the text to be displayed: ", false)

	fmt.Println("\nInput: " + input)
}
