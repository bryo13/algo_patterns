package twopointers

// list has to be sorted
// we go through the list twice hence O(nxn)
func TwoSum(list []int, sum int) bool {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {

			if i == j {
				continue
			}

			if list[i]+list[j] == sum {
				return true
			}

			if list[i]+list[j] > sum {
				break
			}
		}
	}

	return false
}
