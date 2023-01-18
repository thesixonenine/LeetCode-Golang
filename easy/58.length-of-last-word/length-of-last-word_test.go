package length_of_last_word

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testMap := map[string]int{"Hello World": 5, "   fly me   to   the moon  ": 4, "luffy is still joyboy": 6}
	for k, want := range testMap {
		got := lengthOfLastWord(k)
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
}
