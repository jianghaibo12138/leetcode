package main

import (
	"fmt"
)

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
	carry := 0
	var root = &ListNode{}
	var lx = root.Next
	for {
		if l1 != nil && l2 != nil {
			lx.Val = l1.Val + l2.Val + carry
			n := lx.Val - 10
			if n >= 0 {
				carry = 1
				lx.Val = n
			} else {
				carry = 0
			}
			l1 = l1.Next
			l2 = l2.Next
			lx = lx.Next
		} else {
			if l1 == nil {
				if carry != 0 {
					lx.Val = l2.Val + carry
					if lx.Val >= 10 {
						lx.Val -= 10
					}
				}
				break
			}
			if l2 == nil {
				if carry != 0 {
					lx.Val = l1.Val + carry
					if lx.Val >= 10 {
						lx.Val -= 10
					}
				}
				lx = l1.Next
				break
			}
		}
	}
	return root.Next
}

func printF(ln *ListNode) {
	for {
		if ln == nil {
			return
		}
		fmt.Print(ln.Val)
		ln = ln.Next
	}
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
		},
	}
	l3 := addTwoNumbers(&l1, &l2)
	printF(l3)
}
