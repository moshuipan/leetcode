package leetcode

// 将字符串 "PAYPALISHIRING" 以Z字形排列成给定的行数：

// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后从左往右，逐行读取字符："PAHNAPLSIIGYIR"

// 实现一个将字符串进行指定行数变换的函数:

// string convert(string s, int numRows);
// 示例 1:

// 输入: s = "PAYPALISHIRING", numRows = 3
// 输出: "PAHNAPLSIIGYIR"
// 示例 2:

// 输入: s = "PAYPALISHIRING", numRows = 4
// 输出: "PINALSIGYAHRPI"
// 解释:

// P     I    N
// A   L S  I G
// Y A   H R
// P     I

func convert(s string, numRows int) string {
	// 0        8
	// 1      7 9
	// 2    6   10
	// 3  5     11
	// 4        12
	// 整列差	d=numRows*2-2
	// 单列    2*i
	var ns string
	d := numRows*2 - 2 // 整列差
	if d == 0 {
		return s
	}
	l := len(s)
	for i := 0; i < numRows; i++ {
		j := 0 //第几个整列
		for ; ; j++ {
			k := j*d + i //整列的下标
			if k >= l {
				break
			}
			ns += string(s[k])
			if i != 0 && i != numRows-1 {
				k += (numRows - i - 1) * 2 //下一个单列的下标
				if k >= l {
					break
				}
				ns += string(s[k])
			}
		}
	}
	return ns
}
