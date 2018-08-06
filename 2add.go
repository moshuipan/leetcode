package leetcode

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
	l3 := &ListNode{}
	tmp := l3
	x, y := l1, l2
	var add bool
	for {
		var v int
		if add {
			v = x.Val + y.Val + 1
		} else {
			v = x.Val + y.Val
		}
		add = (v / 10) > 0
		tmp.Val = v % 10
		x = x.Next
		y = y.Next
		if x == nil && y == nil {
			tmp.Next = nil
			if add {
				tmp.Next = &ListNode{Val: 1}
			}
			break
		}
		if x == nil && y != nil {
			if add {
				add = false
				x = &ListNode{Val: 1}
			} else {
				tmp.Next = y
				break
			}
		} else if x != nil && y == nil {
			if add {
				add = false
				y = &ListNode{Val: 1}
			} else {
				tmp.Next = x
				break
			}
		}
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	return l3
}
