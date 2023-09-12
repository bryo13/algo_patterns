package sliding_window

func MaxSum(arraylist []int, window int) int {
	var windowSum int = 0
	maxSum := 0

	for i := 0; i < len(arraylist); i++ {
		windowSum += arraylist[i]

		if windowSum > maxSum {
			maxSum = windowSum
		}

		if i >= window-1 {
			windowSum -= arraylist[i-window+1]
		}
	}
	return maxSum
}
