package main

import (
	"fmt"
	"log"

	"yaya.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Toner")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
