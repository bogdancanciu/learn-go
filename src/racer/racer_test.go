package racer

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURL := "http://www.quii.dev"

	want := fastURL
	got := Race(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}