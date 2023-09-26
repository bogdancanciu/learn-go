package hello

import "testing"

func TestGreetWorld(t *testing.T) {
	got := GreetWorld()
	want := "Hello, world!"

	assertCorrectMessage(t, got, want)
}

func TestGreetPerson(t *testing.T) {
	got := GreetPerson("John")
	want := "Hello, John!"

	assertCorrectMessage(t, got, want)
}

func TestEnglishGreetWithArguments(t *testing.T) {
	got := Greet("argument", "English")
	want := "Hello, argument!"

	assertCorrectMessage(t, got, want)
}

func TestEnglishGreetWithEmptyArgument(t *testing.T) {
	got := Greet("", "English")
	want := "Hello, world!"

	assertCorrectMessage(t, got, want)
}

func TestRomanianGreetWithArguments(t *testing.T) {
	got := Greet("argument", "Romanian")
	want := "Salut, argument!"

	assertCorrectMessage(t, got, want)
}

func TestRomanianGreetWithEmptyArgument(t *testing.T) {
	got := Greet("", "English")
	want := "Salut, world!"

	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
