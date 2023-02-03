package add_binary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	a := "11"
	b := "1"
	want := "100"
	got := addBinary(a, b)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}

	a = "1010"
	b = "1011"
	want = "10101"
	got = addBinary(a, b)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}

	a = "0"
	b = "0"
	want = "0"
	got = addBinary(a, b)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
