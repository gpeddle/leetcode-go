package leetcode

func removeDuplicates(nums []int) int {

	pos := 0
	found := -101
	count := 0

	for ii := 0; ii < len(nums); ii++ {
		x := nums[ii]

		if x == found {
			nums[ii] = 9999
		} else if x > found {
			found = x
			nums[pos] = x
			pos++
			count++
		} else {

		}
	}

	return count
}
