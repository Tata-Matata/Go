package main

import (
	"bufio"
	"fmt"
	"helloworld/doctor"
	"os"
	"strings"
)

func main() {

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
