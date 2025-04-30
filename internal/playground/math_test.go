package playground

import "testing"

func TestSum(t *testing.T) {
	got := Sum(2, 3)
	want := 5

	if got != want {
		t.Errorf("Sum(2, 3) = %d; want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := Multiply(4, 5)
	want := 20

	if got != want {
		t.Errorf("Multiply(4, 5) = %d; want %d", got, want)
	}
}
