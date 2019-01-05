package degolomb

import (
	"testing"
)

func TestBitArray(t *testing.T) {
	expectation := []int{0, 0, 1, 0, 0, 1, 1, 1}
	bits := BitArray(byte(39))
	for i, bit := range bits {
		if bit != expectation[i] {
			t.FailNow()
		}
	}
}
func TestDegolomb(t *testing.T) {
	input := byte(39)
	e := Degolomb(input, []int{})
	if len(e) != 1 {
		t.Fatalf("expected a single element, got %d\n", len(e))
	}
	if e[0] != 39 {
		t.Fatalf("expected 39, got %d\n", e[0])
	}

	expectation := []int{0, 1, 7}
	e = Degolomb(input, []int{1, 2, 5})
	if len(e) != len(expectation) {
		t.Fatalf("expected %d elements, got %d\n", len(expectation), len(e))
	}
	for i := range expectation {
		if e[i] != expectation[i] {
			t.Fatalf("%v: expected %d, got %d: \n\texpected %v\n\tactual %v\n", input, expectation[i], e[i], expectation, e)
		}
	}
	input = byte(103)
	expectation = []int{0, 3, 7}
	e = Degolomb(input, []int{1, 2, 5})
	for i := range expectation {
		if e[i] != expectation[i] {
			t.Fatalf("%v: expected %d, got %d: \n\texpected %v\n\tactual %v\n", input, expectation[i], e[i], expectation, e)
		}
	}
}
