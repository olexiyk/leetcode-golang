package leet

/*
 * @lc app=leetcode id=459 lang=golang
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	for i := range s {
		slice := s[:i+1]
		sliceLen := len(slice)
		if len(s)%sliceLen != 0 || len(s)/sliceLen == 1 {
			continue
		}
		repeatTimes := len(s) / sliceLen
		repeated := strings.Repeat(slice, repeatTimes)
		if repeated == s {
			return true
		}
	}

	return false
}

// @lc code=end
