package main

import (
	"fmt"
)

const yoruba = "Yoruba"
const spanish = "Spanish"
const german = "German"
const greeting = "Hello, "
const yorubaGreeting = "Bawo ni, "
const spanishGreeting = "Hola, "
const germanGreeting = "Hallo, "
// domain code: this means this function returns a string
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
	// if language == yoruba {
	// 	return yorubaGreeting + name
	// }
	// if language == spanish {
	// 	return spanishGreeting + name
	// }
	// if language == german {
	// 	return germanGreeting + name
	// }
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case yoruba:
		prefix = yorubaGreeting
	case spanish:
		prefix = spanishGreeting
	case german:
		prefix = germanGreeting
	default:
		prefix = greeting
	}
	return
}

// entry point of the program: exposes the domain code to the outside world
func main() {
	fmt.Println(Hello("world", "German"))
}