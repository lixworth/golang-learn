package main

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("lixworth")
	fmt.Printf(message)
}
