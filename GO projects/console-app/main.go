package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for{
		fmt.Print("->")

		// sometimes functions return more than one value. this underscore means i don't care about the second value that is returned
		userInput, _ := reader.ReadString('\n')
		// userInput = strings.Replace(userInput, "\r", "", -1)
		// userInput = strings.Replace(userInput, "\n", "", -1)

		
		//It removes both \n and \r automatically (and other whitespace):
		userInput = strings.TrimSpace(userInput)
		if userInput == "quit"{
			break
		}else{
			fmt.Println(userInput)
		}

	}
}