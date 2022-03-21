package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// merge into third array, retaining sorted
	m := len(nums1)
	p1 := 0
	n := len(nums2)
	p2 := 0
	combined := m + n
	var merged []int

	for ii := 0; ii < combined; ii++ {
		if p2 == n {
			merged = append(merged, nums1[p1])
			p1++
		} else if p1 == m {
			merged = append(merged, nums2[p2])
			p2++
		} else if nums1[p1] <= nums2[p2] {
			merged = append(merged, nums1[p1])
			p1++
		} else {
			merged = append(merged, nums2[p2])
			p2++
		}
	}

	median := 0.0
	if combined%2 == 1 {
		pos := (combined - 1) / 2
		median = float64(merged[pos])
	} else {
		pos1 := (combined / 2) - 1
		pos2 := combined / 2
		med1 := float64(merged[pos1])
		med2 := float64(merged[pos2])
		median = (med1 + med2) / 2
	}

	return median
}
