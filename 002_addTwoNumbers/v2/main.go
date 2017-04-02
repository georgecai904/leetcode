package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln ListNode) insertNode(entries []int) ListNode {
	var temp *ListNode

	for i, val := range entries {
		if i == 0 {
			ln.Val = val
			temp = &ln
			continue
		} else {
			var next ListNode
			next.Val = val
			temp.Next = &next
			temp = &next
		}
	}
	return ln
}

func (ln ListNode) retrieveListNode() []int {
	var head *ListNode = &ln
	var result []int
	for head != nil {
		result = append(result, head.Val)
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
	}
	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	quotient := 0
	val := 0
	var root ListNode
	var temp *ListNode
	// Assign address of root to temp, then action on temp will directly change on root, but will not change address of root
	temp = &root
	for l1 != nil {
		val = (l1.Val + l2.Val + quotient) % 10
		quotient = (l1.Val + l2.Val) / 10
		temp.Val = val
		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil {
			var next ListNode
			temp.Next = &next
			temp = &next
		}
	}
	return &root
}

func main() {
	var l1, l2 ListNode
	l1 = l1.insertNode([]int{2, 4, 3})
	l2 = l2.insertNode([]int{5, 6, 4})
	fmt.Println(addTwoNumbers(&l1, &l2).retrieveListNode())
}
