package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 8)
	expected := "aaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}