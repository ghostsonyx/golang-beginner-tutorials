package main

import (
	"fmt"
	"log"

	"HOME-GO/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Sarah", "Bob", "William"}

	messages, err := greetings.MultipleGreetings(names)

	message, err := greetings.SingleGreeting("Nathan")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	fmt.Println(message)
}
