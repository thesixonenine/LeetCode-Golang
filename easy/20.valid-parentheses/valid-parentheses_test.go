package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	testMap := map[string]bool{}
	testMap["()"] = true
	testMap["()[]{}"] = true
	testMap["(]"] = false
	testMap["){"] = false

	for s, want := range testMap {
		got := isValid(s)
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
}
