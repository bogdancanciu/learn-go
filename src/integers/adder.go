package integers

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

func SumAll(numbersToSum ...[]int) []int {
	numbersLength := len(numbersToSum)
	sums := make([]int, numbersLength)

	for i, numbers:= range numbersToSum {
		sums[i] = Sum(numbers)
	}
	fmt.Println(sums)
	return sums
}