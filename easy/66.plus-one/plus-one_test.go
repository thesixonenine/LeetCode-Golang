package plus_one

import "testing"

func TestPlusOne(t *testing.T) {
	testArr := []int{1, 2, 3}
	wantArr := []int{1, 2, 4}
	got := plusOne(testArr)
	if len(wantArr) != len(got) {
		t.Errorf("got length '%v' want length '%v'", len(got), len(wantArr))
	}
	for i, item := range wantArr {
		if item != got[i] {
			t.Errorf("got '%v' want '%v'", got[i], item)
		}
	}
}
