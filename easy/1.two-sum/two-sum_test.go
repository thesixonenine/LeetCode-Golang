package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	// 示例1
	numbs := []int{2, 7, 11, 15}
	target := 9
	got := twoSum(numbs, target)
	want := []int{0, 1}
	arrayEqual(got, want, t)

	// 示例2
	numbs = []int{3, 2, 4}
	target = 6
	got = twoSum(numbs, target)
	want = []int{1, 2}
	arrayEqual(got, want, t)

	// 示例3
	numbs = []int{3, 3}
	target = 6
	got = twoSum(numbs, target)
	want = []int{0, 1}
	arrayEqual(got, want, t)
}
func arrayEqual(got []int, want []int, t *testing.T) {
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	}
}
