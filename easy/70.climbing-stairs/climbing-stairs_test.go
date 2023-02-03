package climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	testMap := map[int]int{
		2: 2, 3: 3, 45: 1836311903,
	}

	for k, want := range testMap {
		got := climbStairs(k)
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
}
