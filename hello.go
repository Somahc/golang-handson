package main

import (
	"fmt"
	"go-handson/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!")
}
