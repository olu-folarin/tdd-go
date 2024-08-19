package main

import (
	"fmt"
)

const greeting = "Hello, "
// domain code: this means this function returns a string
func Hello(name string) string {
	return greeting + name
}

// entry point of the program: exposes the domain code to the outside world
func main() {
	fmt.Println(Hello("world"))
}