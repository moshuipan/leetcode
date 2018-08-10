package leetcode

// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为1000。
// 示例 1：
// 输入: "babad"
// 输出: "bab"
// 注意: "aba"也是一个有效答案。
// 示例 2：
// 输入: "cbbd"
// 输出: "bb"
func longestPalindrome(s string) string {
	var maxs string
	var maxl int
	sl := len(s)
	for i := 0; i < sl; i++ {
		j := 1
		for i-j >= 0 && i+j < sl { // s[i] 的左右相等
			if s[i-j] == s[i+j] {
				j++
			} else {
				break
			}
		}
		j -= 1
		if 2*j+1 > maxl {
			maxs = s[i-j : i+j+1]
			maxl = 2*j + 1
		}
		j = 0
		for i-j >= 0 && i+j+1 < sl { // s[i]==s[i+1]
			if s[i-j] == s[i+j+1] {
				j++
			} else {
				break
			}
		}
		if 2*j > maxl { // 为1 不考虑
			maxs = s[i-j+1 : i+j+1]
			maxl = 2 * j
		}
	}
	return maxs
}
