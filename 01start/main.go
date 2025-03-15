package main

import (
	"bufio"
	"fmt"
	"helloworld/doctor"
	"math/rand"
	"os"
	"strings"
)

const prompt = " and press ENTER"

func main() {
	guessingGame()
}

func guessingGame() {
	random := rand.New(rand.NewSource(123))

	firstNum := random.Intn(8) + 2
	secondNum := random.Intn(8) + 2
	subtraction := random.Intn(8) + 2
	answer := firstNum*secondNum - subtraction

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Let's play the guessing game. Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by ", firstNum, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNum, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	fmt.Println("And the answer is ", answer)

}

func chatbotElisa() {
	fmt.Println("Welcome to the chatbot")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("->")
		userInput, _ := reader.ReadString('\n')

		//newline for linux
		trimmedInput := strings.Replace(userInput, "\n", "", -1)

		//newline for windows
		trimmedInput = strings.Replace(trimmedInput, "\r\n", "", -1)

		if trimmedInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}
}

func demoVariableAssignment() {

	//assign without type
	message := "hello world"
	sayHello(message)

	//assign later
	var messageTwo string
	messageTwo = "hello again"
	fmt.Println(messageTwo)

	//assign inline
	var messageThree string = "and hello one more time"
	fmt.Println(messageThree)
}

func sayHello(message string) {

	fmt.Println(message)
}
