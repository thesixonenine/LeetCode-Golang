package palindrome_number

import "testing"

func TestPalindromeNumber(t *testing.T) {
	intMap := map[int]bool{}
	intMap[121] = true
	intMap[-121] = false
	intMap[10] = false
	intMap[21120] = false

	for k, want := range intMap {
		palindrome(k, want, t)
	}
}

func palindrome(k int, want bool, t *testing.T) {
	got := isPalindrome(k)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
	got = isPalindrome1(k)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
