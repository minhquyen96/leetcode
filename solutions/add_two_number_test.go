package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumber(t *testing.T) {
	e1 := ListNode{
		Val: 3,
	}
	e2 := ListNode{
		Val:  4,
		Next: &e1,
	}
	e3 := ListNode{
		Val:  2,
		Next: &e2,
	}

	e4 := ListNode{
		Val: 4,
	}
	e5 := ListNode{
		Val:  6,
		Next: &e4,
	}
	e6 := ListNode{
		Val:  5,
		Next: &e5,
	}

	rs := addTwoNumbers(&e3, &e6)
	fmt.Println(rs)
}

func TestAddTwoNumber1(t *testing.T) {
	e1 := ListNode{
		Val: 0,
	}

	e2 := ListNode{
		Val: 0,
	}

	rs := addTwoNumbers(&e1, &e2)
	fmt.Println(rs)
}

func TestAddTwoNumber3(t *testing.T) {
	e1 := ListNode{
		Val: 9,
	}
	e2 := ListNode{
		Val:  9,
		Next: &e1,
	}
	e3 := ListNode{
		Val:  9,
		Next: &e2,
	}

	e4 := ListNode{
		Val:  9,
		Next: &e3,
	}

	e5 := ListNode{
		Val:  9,
		Next: &e4,
	}

	e6 := ListNode{
		Val:  9,
		Next: &e5,
	}

	f1 := ListNode{
		Val: 9,
	}
	f2 := ListNode{
		Val:  9,
		Next: &f1,
	}
	f3 := ListNode{
		Val:  9,
		Next: &f2,
	}
	f4 := ListNode{
		Val:  9,
		Next: &f3,
	}
	rs := addTwoNumbers(&e6, &f4)
	fmt.Println(rs)
}
