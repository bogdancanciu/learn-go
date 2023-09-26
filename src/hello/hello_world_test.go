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