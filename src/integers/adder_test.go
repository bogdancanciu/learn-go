package integers

import (
	"testing"
	"reflect"
)

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
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2,3}, []int{9,5,1})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}
}