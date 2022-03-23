package leetcode

func isMatch(s string, p string) bool {

	const DOT = uint8('.')  //46
	const STAR = uint8('*') //42

	if len(p) > len(s) {
		return false
	}

	//scan pattern for regex chars
	var pattern []uint8
	for pp := 0; pp < len(p); pp++ {
		c := p[pp]
		if c == DOT {
			pattern = append(pattern, c)
		}
		if c == STAR {
			pattern = append(pattern, c)
		}
	}

	return false
}
