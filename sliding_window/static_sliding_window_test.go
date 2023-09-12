package sliding_window

import "testing"

func TestMaxSum(t *testing.T) {
	arrayInput := []int{10, 20, 10, 30, 5}
	want := 60
	maxSum := MaxSum(arrayInput, 3)

	if want != maxSum {
		t.Errorf("Want %d but got %d", want, maxSum)
	}
}
