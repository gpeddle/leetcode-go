package leetcode

func getConcatenation(nums []int) []int {
	var ary []int
	ary = append(ary, nums...)
	ary = append(ary, nums...)
	return ary
}

/*
// simple make and iterate approach:
func getConcatenation(nums []int) []int {

	ln := len(nums)
	ary := make([]int, ln*2)

	for ii := 0; ii < ln; ii++ {
		ary[ii] = nums[ii]
		ary[ii+ln] = nums[ii]
	}
	return ary
}
*/
