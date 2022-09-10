package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	rs := ListNode{}
	rs.Val = l1.Val + l2.Val
	if rs.Val > 9 {
		rs.Val -= 10
		if l1.Next != nil {
			l1.Next.Val++
		} else if l2.Next != nil {
			l2.Next.Val++
		} else {
			l1.Next = &ListNode{
				Val: 1,
			}
		}
	}
	if l1.Next != nil && l2.Next != nil {
		rs.Next = addTwoNumbers(l1.Next, l2.Next)
	} else if l1.Next == nil && l2.Next != nil {
		rs.Next = addTwoNumbers(&ListNode{}, l2.Next)
	} else if l1.Next != nil && l2.Next == nil {
		rs.Next = addTwoNumbers(l1.Next, &ListNode{})
	}
	return &rs
}

type ListNode struct {
	Val  int
	Next *ListNode
}
