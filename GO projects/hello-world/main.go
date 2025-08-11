package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		//for windows
		userInput = strings.Replace(userInput, "\r\n", "", -1)

		// for mac, linux , etc.
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {

			fmt.Println(doctor.Response(userInput))

		}
	}
}
