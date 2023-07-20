package main

import (
	"fmt"

	"github.com/allancapistrano/user-input"
)

func main() {
	input := userinput.GetUserInput("Digite o texto para ser exibido: ")

	fmt.Println("\nO usuário digitou: " + input)
}
