package sliding_window

func MaxSum(arraylist []int, window int) int {
	// windowSum holds the total sum of the
	// numbers within the window
	var windowSum int = 0

	// maxSum is the greatest sum within the windows
	maxSum := 0

	for i := 0; i < len(arraylist); i++ {
		// we initialize the window sum with the first element in
		// the array
		windowSum += arraylist[i]

		if windowSum > maxSum {
			maxSum = windowSum
		}

		// removes the first element in the window when the window width is reached
		if i >= window-1 {
			windowSum -= arraylist[i-window+1]
		}
	}
	return maxSum
}
