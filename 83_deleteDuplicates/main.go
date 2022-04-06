package main

import "fmt"

//
type ListNode struct {
	Val   int
	Index int
	Next  *ListNode
}

func main() {

	var l1 ListNode
	var l2 ListNode
	var l3 ListNode
	var l4 ListNode
	var l5 ListNode
	var l6 ListNode

	l1.Val = 1
	l1.Index = 0
	l2.Val = 1
	l2.Index = 1
	l3.Val = 1
	l3.Index = 2
	l4.Val = 2
	l4.Index = 3
	l5.Val = 3
	l5.Index = 4
	l6.Val = 3
	l6.Index = 5

	l1.Next = &l2
	l2.Next = &l3
	//l3.Next = &l4
	//l4.Next = &l5
	//l5.Next = &l6

	//for i := 1; i < 5; i++ {
	//	var node ListNode
	//	node.Val = i / 2
	//	tmp.Next = &node
	//	tmp = tmp.Next
	//}

	result := deleteDuplicates(&l1)
	fmt.Println(result)
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil && node.Next != nil {
		if node.Val == node.Next.Val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	return head
}
