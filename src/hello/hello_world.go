package hello

import "fmt"

const greetingPrefix = "Hello,"

func GreetWorld() string {
	return greetingPrefix + " world!"
}

func GreetPerson(personName string) string {
	return fmt.Sprintf("%s %s!", greetingPrefix, personName)
}
