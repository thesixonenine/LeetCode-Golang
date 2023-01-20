package plus_one

import "testing"

func TestPlusOne(t *testing.T) {
	testArr := []int{1, 2, 3}
	wantArr := []int{1, 2, 4}
	got := plusOne(testArr)
	check(t, wantArr, got)

	testArr = []int{4, 3, 2, 1}
	wantArr = []int{4, 3, 2, 2}
	got = plusOne(testArr)
	check(t, wantArr, got)

	testArr = []int{0}
	wantArr = []int{1}
	got = plusOne(testArr)
	check(t, wantArr, got)

	testArr = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	wantArr = []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 7}
	got = plusOne(testArr)
	check(t, wantArr, got)
}

func check(t *testing.T, wantArr []int, got []int) {
	if len(wantArr) != len(got) {
		t.Errorf("got length '%v' want length '%v'", len(got), len(wantArr))
		return
	}
	for i, item := range wantArr {
		if item != got[i] {
			t.Errorf("got '%v' want '%v'", got[i], item)
		}
	}
}
