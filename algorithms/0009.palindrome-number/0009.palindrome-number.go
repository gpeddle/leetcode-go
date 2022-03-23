package leetcode

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	var p1 int
	var p2 int

	s := fmt.Sprint(x)
	length := len(s)

	if length%2 == 1 {
		p1 = ((length - 1) / 2) - 1
		p2 = p1 + 2
	} else {
		p1 = (length / 2) - 1
		p2 = p1 + 1
	}

	for p1 >= 0 {
		if s[p1] != s[p2] {
			return false
		}
		p1--
		p2++
	}

	return true
}
