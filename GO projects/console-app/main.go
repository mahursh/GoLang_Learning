package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	fmt.Println("Press any key on the keyboard. Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyEsc {
			break
		}

		// Handle printable keys
		if key == 0 {
			fmt.Printf("You pressed: %q\n", char)
		} else {
			fmt.Printf("You pressed special key: %v\n", key)
		}
	}

	fmt.Println("Program exiting...")
}
