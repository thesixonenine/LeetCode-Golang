package roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	intMap := map[string]int{}
	intMap["III"] = 3
	intMap["IV"] = 4
	intMap["IX"] = 9
	intMap["LVIII"] = 58

	for roman, want := range intMap {
		got := romanToInt(roman)
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
}
