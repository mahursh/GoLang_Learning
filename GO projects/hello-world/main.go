package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

func main() {
	reader  := bufio.NewReader(os.Stdin)
	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)
	for{
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		response := doctor.Response(userInput)
		fmt.Println(response)
		}
	}

