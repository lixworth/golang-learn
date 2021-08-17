package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	message1, err := greetings.Hello("rumao")
	message, err := greetings.Hello("233")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(message1)
	fmt.Printf(message)
}
