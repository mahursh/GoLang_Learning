package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwensADog      bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User

	user.UserName = readString("What is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorit number?")
	user.OwensADog = GetYesOrNo("Do you own a dog?")

	fmt.Printf("Your name is %s ,and you are %d years old . your faviorit  numbrt is %.2f . %s\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		map[bool]string{true: "you do own a dog", false: " you don,t have a dog."}[user.OwensADog])

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

func GetYesOrNo(q string) bool {
	fmt.Println(q)
	prompt()
	err := keyboard.Open()
	if err != nil {
		fmt.Println("an error accured")
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(q)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Please enter y for y or n for no.")
		}
		if char == 'n' || char == 'N' {
			return false
		}
		return true
	}
}
