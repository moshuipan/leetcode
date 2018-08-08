package leetcode

// 给定一个字符串，找出不含有重复字符的最长子串的长度。

// 示例：

// 给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

// 给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

// 给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
func lengthOfLongestSubstring(s string) int {
	tmp := make(map[rune]int)
	var maxL, currl, z int
	for i, ss := range s {
		if v, ok := tmp[ss]; !ok {
			tmp[ss] = i
			currl++
			if currl > maxL {
				maxL = currl
			}
		} else {
			if v < z && z < i {
				v = z
			}
			currl = i - v
			tmp[ss] = i
			if currl > maxL {
				maxL = currl
			}
			z = v
		}
	}
	return maxL
}
