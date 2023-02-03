package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	testMap := map[int]int{
		4: 2, 8: 2, 16: 4, 65: 8, 0: 0,
	}

	for k, want := range testMap {
		got := mySqrt(k)
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
}
