package total

func Total(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

