package hello

import "fmt"

const (
	romanian = "Romanian"

	emptySpace = " "
	englishGreetingPrefix = "Hello,"
	romanianGreetingPrefix = "Salut,"
)

func GreetWorld() string {
	return englishGreetingPrefix + " world!"
}

func GreetPerson(personName string) string {
	return fmt.Sprintf("%s %s!", englishGreetingPrefix, personName)
}

func Greet(greeted, language string) string {
	if greeted == "" {
		greeted = "world"
	}

	return greetingPrefix(language) + emptySpace + greeted + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case romanian:
		prefix = romanianGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}

	return
}
