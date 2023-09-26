package hello

import "testing"

func TestGreetWorld(t *testing.T) {
	got := GreetWorld()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreetPerson(t *testing.T) {
	got := GreetPerson("John")
	want := "Hello, John!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreetWithArguments(t *testing.T) {
	got := Greet("argument")
	want := "Hello, argument!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreetWithEmptyArgument(t *testing.T) {
	got := Greet("")
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
