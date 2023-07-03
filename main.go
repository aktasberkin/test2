package main

import (
	"fmt"

	"github.com/emiellecarde/test1/hello"
	"github.com/emiellecarde/test1/no"
)

func main() {
	fmt.Println("Test")
	hello.Greetings()
	no.No()
	fmt.Println("Checkpoint")
	fmt.Println("Checkpoint2")
	fmt.Println("Checkpoint3")
}
