package leetcode

func twoSum(nums []int, target int) []int {

	visited := make(map[int]int)
	for currentIndex, currentValue := range nums {
		if candidateIndex, found := visited[target-currentValue]; found {
			return []int{candidateIndex, currentIndex}
		}
		visited[currentValue] = currentIndex
	}

	return []int{}
}
