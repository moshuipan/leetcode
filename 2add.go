package leetcode

import (
	"fmt"
	"math"
)

// 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。

// 示例：

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	v1 := 0
	for i := 1; ; {
		if l1 == nil {
			break
		}
		v1 += l1.Val * i
		i = i * 10
		l1 = l1.Next
	}
	for i := 1; ; {
		if l2 == nil {
			break
		}
		v1 += l2.Val * i
		i = i * 10
		l2 = l2.Next
	}
	l := len(fmt.Sprint(v1))
	l3 := &ListNode{}
	tmp := l3
	for i := 1; i <= l; i++ {
		tmp.Val = v1 % int(math.Pow10(i))
		v1 = (v1 - tmp.Val) / int(math.Pow10(i))
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	return l3
}
