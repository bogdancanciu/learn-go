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

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	defaultValue := 0

	for _, numbers:= range numbersToSum {
		if hasElements(numbers){
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, defaultValue)
		}
	}

	return
}

func hasElements(numbers []int) bool {
	return len(numbers) > 0
}