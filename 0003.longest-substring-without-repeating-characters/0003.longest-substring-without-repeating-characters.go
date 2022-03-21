package leetcode

func lengthOfLongestSubstring(s string) int {

	longest := 0

	for ii := 0; ii < len(s); ii++ {
		letters := make(map[uint8]int)
		for jj := ii; jj < len(s); jj++ {
			cand := s[jj]
			_, dupe := letters[cand]

			if dupe {
				//reset
				break
			}
			letters[cand] = jj
		}
		if longest == 0 || longest < len(letters) {
			longest = len(letters)
		}
	}
	return longest
}

/*
func lengthOfLongestSubstring(s string) int {

	found := make(map[int32]bool)
	longest := 0

	for _, letter := range s {
		_, present := found[letter]
		if present {
			//reset
			if len(found) > longest {
				longest = len(found)
			}
			found = make(map[int32]bool)
			found[letter] = true
		} else {
			found[letter] = true
		}
	}
	if len(found) > longest {
		longest = len(found)
	}
	if longest == 0 {
		longest = len(found)
	}

	return longest
}
*/
