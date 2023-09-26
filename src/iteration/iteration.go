package iteration


func Repeat(character string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += character
	}

	return
}