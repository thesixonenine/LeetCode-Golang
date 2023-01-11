package roman_to_integer

import "testing"

func TestRomanToInt(t *testing.T) {
	testMap := map[string]int{}
	testMap["III"] = 3
	testMap["IV"] = 4
	testMap["IX"] = 9
	testMap["LVIII"] = 58

	for roman, want := range testMap {
		got := romanToInt(roman)
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
}
