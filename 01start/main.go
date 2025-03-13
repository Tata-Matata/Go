package main

import "fmt"

func main() {

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
