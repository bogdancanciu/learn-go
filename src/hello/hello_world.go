package hello

import "fmt"

const greetingPrefix = "Hello,"

func GreetWorld() string {
	return greetingPrefix + " world!"
}

func GreetPerson(personName string) string {
	return fmt.Sprintf("%s %s!", greetingPrefix, personName)
}

func Greet(greeted string) string {
	result := GreetPerson(greeted)

	if greeted == "" {
		result = GreetWorld()
	}

	return result
}
