package leetcode

import (
	"strings"
)

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

// Manacher算法
//     ----------P0--r--j--r-----M-----r--i--r--P1---------------
//     --------r--P0----j-----r--M--r-----i----P1--r-------------
//     ----------P0-----------M-----------P1---i------------
func longestPalindromeByManacher(s string) string {
	ns := "#"
	for i := 0; i < len(s); i++ { // 初始化字符串,字符间加#
		ns += string(s[i]) + "#"
	}
	l := len(ns)
	pl := make([]int, l) // 最大回文串半径
	maxI := 0            // 半径最大的回文串中心值下标
	maxS := ""           // 最大子字符串
	pl[0] = 1
	for i := 1; i < l; i++ {
		p := maxI + pl[maxI] - 1
		if i >= p { // 依次判断回文
			j := 1
			for i-j >= 0 && i+j < l {
				if ns[i-j] == ns[i+j] {
					j++
				} else {
					break
				}
			}
			j--
			pl[i] = j + 1
			// if pl[i] > 2 {
			maxI = i
			// }
			if 2*pl[i]-1 > len(maxS) {
				maxS = ns[i-j : i+j+1]
			}
		} else { // 找对称点
			symmetryP := 2*maxI - i
			if pl[symmetryP]-1 < p-i { // pl[i]=pl[symmetryP]
				pl[i] = pl[symmetryP]
			} else { // x>i+pl[symmetryP]-1,依次比对
				j := p - i + 1
				for i-j >= 0 && i+j < l {
					if ns[i-j] == ns[i+j] {
						j++
					} else {
						break
					}
				}
				pl[i] = j
				// if pl[i] > 2 {
				maxI = i
				// }
				if 2*pl[i]-1 > len(maxS) {
					maxS = ns[i-pl[i]+1 : i+pl[i]]
				}
			}
		}
	}
	return strings.Replace(maxS, "#", "", -1)
}
