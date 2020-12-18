package main

import (
	"fmt"
	"log"

	"yaya.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Toner",
		"Emrys",
		"Frank",
	}

	messages, err := greetings.Hellos(names)

	// message, err := greetings.Hello("Toner")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	// fmt.Println(message)
}
