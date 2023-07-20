package userinput

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

func GetUserInput(displayMessage string) string {
	if err := keyboard.Open(); err != nil {
		log.Fatal("Unable to use the keyboard.")
	}
	defer keyboard.Close()

	fmt.Print(displayMessage)
	
	var input string
	var cursorIndex int

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			os.Exit(0)
		}

		if key == keyboard.KeyEnter {
			break
		}

		if key == keyboard.KeyBackspace || key == keyboard.KeyBackspace2 {
			if cursorIndex > 0 {
				input = input[:cursorIndex-1] + input[cursorIndex:]
				cursorIndex--
			}
		} else if key == keyboard.KeyArrowLeft {
			if cursorIndex > 0 {
				cursorIndex--
			}
		} else if key == keyboard.KeyArrowRight {
			if cursorIndex < len(input) {
				cursorIndex++
			}
		} else if key == keyboard.KeyHome {
			if cursorIndex >= 0 && cursorIndex <= len(input) {
				cursorIndex = 0
			}
		} else if key == keyboard.KeyEnd {
			if cursorIndex >= 0 && cursorIndex <= len(input) {
				cursorIndex = len(input)
			}
		} else if key == keyboard.KeySpace {
			input = input[:cursorIndex] + " " + input[cursorIndex:]
			cursorIndex++
		} else if char != 0 {
			input = input[:cursorIndex] + string(char) + input[cursorIndex:]
			cursorIndex++
		}

		clearLine()
		fmt.Print("\r" + displayMessage + input + " ")
		moveCursorBack(len(input) - cursorIndex + 1)
	}

	return input
}

func clearLine() {
	fmt.Print("\r" + strings.Repeat(" ", 60) + "\r")
}

func moveCursorBack(steps int) {
	if steps > 0 {
		fmt.Print("\033[" + fmt.Sprintf("%d", steps) + "D")
	}
}
