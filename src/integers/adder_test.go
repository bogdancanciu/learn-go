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

	assertSumsAreEqual(t, got, want)
}

func TestSumAllTailsWithPopulatedSlices(t *testing.T) {
	got := SumAllTails([]int{1,2,3}, []int{1,0,0,9}, []int{100,0,1})
	want := []int{5,9,1}

	assertSumsAreEqual(t, got, want)
}

func TestSumAllTailsWithEmptySlice(t *testing.T) {
	got := SumAllTails([]int{}, []int{1})
	want := []int{0,0}

	assertSumsAreEqual(t, got, want)
}

func assertSumsAreEqual(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}
}
