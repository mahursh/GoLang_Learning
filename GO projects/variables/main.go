package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

// constants never change
const prompt = " don't type your number in, just press ENTER when ready."

func main() {
	// seed the random number generator.
	
	var firstNumber = rand.IntN(8) + 2
	var secondNumber = rand.IntN(8) + 2
	var subtraction = rand.IntN(8) + 2
	var answer int //becuase we did not gaved it a value, it will be asiigned with default value , for example for strings , its just a empty string.
	reader := bufio.NewReader(os.Stdin)
	// display a welcom/instructions
	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10 ", prompt)
	reader.ReadString('\n')

	//tgake them, through the game
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you orginialt thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now Subtract", subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer
	answer = firstNumber * secondNumber - subtraction
	fmt.Println("The answer is", answer)

}
