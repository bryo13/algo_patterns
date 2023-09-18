package twopointers

import "testing"

func TestTwoSum(t *testing.T) {
	want := true
	got := TwoSum([]int{1, 2, 3}, 5)

	if want != got {
		t.Errorf("Got %v but wanted %v\n", got, want)
	}
}
