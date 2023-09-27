package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{2,2,4,6,8}
	got := Sum(numbers)
	want := 22

	if got != want {
		t.Errorf("expected '%d' but got '%d'", got, want)
	}
}