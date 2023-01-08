package palindrome_number

import "testing"

func TestPalindromeNumber(t *testing.T) {
	intMap := map[int]bool{}
	intMap[121] = true
	intMap[-121] = false
	intMap[10] = false

	for k, want := range intMap {
		isPalindrome(k, want, t)
	}
}

func isPalindrome(k int, want bool, t *testing.T) {
	got := PalindromeNumber1(k)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
	got = PalindromeNumber2(k)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
