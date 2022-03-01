package main

import (
	"fmt"
	"log"

	"minzhi.io/greetings"
)

func main() {
	log.SetPrefix("greetings:\n")
	log.SetFlags(0)

	names := []string{"minzhi", "guanchen", "yanliu"}

	if messages, err := greetings.Hellos(names); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(messages)
	}
}
