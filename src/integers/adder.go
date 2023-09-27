package integers

func Add(x, y int) int {
	return x + y
}

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}