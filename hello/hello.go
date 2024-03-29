package main

import (
	"fmt"
	"log"

	"sm.com/greetings"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Roman", "Rajiv", "Ravi"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
