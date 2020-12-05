package main

import (
	"fmt"

	"yaya.com/greetings"
)

func main() {
	message := greetings.Hello("Toner")
	fmt.Println(message)
}
