package main

import (
	"fmt"
	"hammer-bitcoin/game"
)

func main() {
	playAgain := true

	for playAgain {
		game.Play() 
		playAgain = game.GetYesOrNo("Would you like to play again (Y/N) ?")
	}

	fmt.Println("")
	fmt.Println("Goodbye.")
}