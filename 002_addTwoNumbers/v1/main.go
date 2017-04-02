package main

import (
	"fmt"
)

//You are given two non-empty linked lists representing two non-negative integers.
//The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
//You may assume the two numbers do not contain any leading zero, except the number 0 itself.

//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8

type ListNode struct {
	val  int
	next *ListNode
}

func (ln ListNode) insertNode(entries []int) ListNode {
	var temp *ListNode

	for i, val := range entries {
		if i == 0 {
			ln.val = val
			temp = &ln
			continue
		} else {
			var next ListNode
			next.val = val
			temp.next = &next
			temp = &next
		}
	}
	return ln
}

func (ln ListNode) retrieveListNode() []int {
	var head *ListNode = &ln
	var result []int
	for head != nil {
		result = append(result, head.val)
		if head.next != nil {
			head = head.next
		} else {
			break
		}
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 ListNode
	var entries []int
	quotient := 0
	for {
		entries = append(entries, (l1.val+l2.val+quotient)%10)
		quotient = (l1.val + l2.val) / 10
		if l1.next == nil {
			break
		} else {
			l1 = l1.next
			l2 = l2.next
		}
	}
	l3 = l3.insertNode(entries)
	return &l3
}

func main() {
	//addTwoNumbers()
	var l1, l2 ListNode
	//a.val = 10
	l1 = l1.insertNode([]int{2, 4, 3})
	l2 = l2.insertNode([]int{5, 6, 4})
	//fmt.Println(l1.retrieveListNode())
	//fmt.Println(l2.retrieveListNode())
	fmt.Println(addTwoNumbers(&l1, &l2).retrieveListNode())
	//
	//b := []int{2,4,3}
	//c := []int{5,6,4}
	//e := make([]int, len(b))
	//fmt.Println(e)
	//for i,_ := range b{
	//	e[i] = (b[i] + c[i] + e[i]) % 10
	//	if i < len(e) - 1{
	//		e[i+1] = (b[i] + c[i]) / 10
	//	}
	//}
	//fmt.Println(e)

}
