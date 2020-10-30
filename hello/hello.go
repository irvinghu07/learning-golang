package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Irving",
		"Beibei",
		"Pipi",
	}
	message, err := greetings.Hellos(names)
	if nil != err{
		log.Fatal(err)
	}
	fmt.Println(message)
}
