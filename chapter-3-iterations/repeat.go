package repeat

import "fmt"

func Repeat(character string, repeater int) (string, int) {
	var repeated string
	for i := 0; i < repeater; i++ {
		repeated += character
	}
	fmt.Println(repeated, len(repeated))
	return repeated, len(repeated)
}

func Adder(alph string) string {
	var cont string
	for i := 0; i < 7; i++ {
		cont += alph
	}

	return cont
}