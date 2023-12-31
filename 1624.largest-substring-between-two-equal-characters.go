package leet

/*
 * @lc app=leetcode id=1624 lang=golang
 *
 * [1624] Largest Substring Between Two Equal Characters
 */

// @lc code=start
func maxLengthBetweenEqualCharacters(s string) int {
	result := -1
	l := make(map[byte]int)
	for i := range s {
		c := s[i]
		_, ok := l[c]

		if !ok {
			l[c] = i
		} else {
			result = max(result, i-l[c]-1)
			result = max(result, 0)
		}
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
