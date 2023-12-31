package leet

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	sum1 := make(map[rune]int)
	sum2 := make(map[rune]int)
	for _, v := range s {
		sum1[v] += 1
	}
	for _, v := range t {
		sum2[v] += 1
	}
	if len(sum1) != len(sum2) {
		return false
	}
	for k, v := range sum1 {
		if sum2[k] != v {
			return false
		}
	}

	return true
}

// @lc code=end
