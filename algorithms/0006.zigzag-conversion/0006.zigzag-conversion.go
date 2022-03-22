package leetcode

import "strings"

func convert(s string, numRows int) string {

	if len(s) <= 1 {
		return s
	}
	if numRows <= 1 {
		return s
	}

	row_data := make([][]uint8, numRows)

	step := 0
	rowIndex := 0
	for pos := 0; pos < len(s); pos++ {

		c := s[pos]
		rowIndex += step
		row := row_data[rowIndex]
		row = append(row, uint8(c))
		row_data[rowIndex] = row

		if rowIndex == 0 {
			step = 1
		} else if rowIndex == numRows-1 {
			step = -1
		}

	}
	var qq []string

	for nn := 0; nn < numRows; nn++ {
		rs := string(row_data[nn])
		qq = append(qq, rs)
	}
	result := strings.Join(qq, "")
	return result
}
