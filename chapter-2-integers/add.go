package integers

import "fmt"

func Add(x, y int) int {
	return x + y
}

func TryAdd() {
	adder := Add(21, 14)
	fmt.Println(adder)
}