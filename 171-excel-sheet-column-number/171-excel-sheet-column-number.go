package leetcode

import "fmt"

// Problem: 171. Excel Sheet Column Number
// Link: https://leetcode.com/problems/excel-sheet-column-number/description/
// Lessons Lerned:
//   - I understand that we have base-26 numiration system, but don't khow about Bijective numeration
//     https://en.wikipedia.org/wiki/Bijective_numeration
func titleToNumber(columnTitle string) int {
	// A - 1, Z - 26, AA - 27, AZ - 52 = base-26 numeration system
	res := 0
	for k, c := range columnTitle {
		fmt.Println(k, c)
		res = res*26 + int(c-'A'+1)
	}

	return res
}
