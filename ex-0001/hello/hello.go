package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Lucas")
	fmt.Println(message)
}
