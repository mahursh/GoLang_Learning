package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwensADog bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorit number?")

	fmt.Printf("Your name is %s ,and you are %d years old . your faviorit  numbrt is %.2f .\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber)

}

func prompt() {
	fmt.Println("->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput != "" {
			return userInput
		} else {
			fmt.Println("Please enter a value.")
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput) // cleaner than Replace
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number.")
			continue // stay in loop until valid
		}
		return num // only return if conversion succeeds
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number.")
		} else {
			return num
		}
	}
}
