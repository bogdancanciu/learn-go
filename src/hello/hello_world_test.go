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

func TestGreetWithArguments(t *testing.T) {
	got := Greet("argument")
	want := "Hello, argument!"

	assertCorrectMessage(t, got, want)
}

func TestGreetWithEmptyArgument(t *testing.T) {
	got := Greet("")
	want := "Hello, world!"

	assertCorrectMessage(t, got, want)
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
