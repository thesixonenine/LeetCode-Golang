package palindrome_number

import "testing"

func TestPalindromeNumber(t *testing.T) {
	testMap := map[int]bool{}
	testMap[121] = true
	testMap[-121] = false
	testMap[10] = false
	testMap[21120] = false

	for k, want := range testMap {
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
