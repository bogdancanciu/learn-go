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

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers:= range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}