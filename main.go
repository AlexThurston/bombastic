package main

import (
	"fmt"
	"github.com/alexthurston/bombastic/lib"
)

func main() {
	a := lib.Foo{"Alex", "Thurston"}
	fmt.Println("Welcome to Bombastic")
	fmt.Println(a.Firstifier())
}
