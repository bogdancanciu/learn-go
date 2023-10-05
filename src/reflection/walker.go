package walker

func walk(x interface{}, fn func(input string)) {
	fn("Some input")
}