package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"firstKey": "this is value"}

	got := dictionary.Search("firstKey")
	want := "this is value"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}