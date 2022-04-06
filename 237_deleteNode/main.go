package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	nTmp1 := &l1
	for i := 0; i < 3; i++ {
		var node ListNode
		node.Val = i + 1
		nTmp1.Next = &node
		nTmp1 = &node
	}

	fmt.Println(l1.Next.Val)
	//deleteNode(&l1)
	//deleteNode(l1.Next)
	deleteNode(l1.Next.Next)

	fmt.Printf("%+v", l1)
}

// 非最优解
func deleteNode(node *ListNode) {
	for {
		if node.Next.Next == nil {
			node.Val = node.Next.Val
			node.Next = nil
			break
		}

		node.Val = node.Next.Val
		node = node.Next
	}
}

// 最优解
//func deleteNode(node *ListNode) {
//	node.Val = node.Next.Val
//	node.Next = node.Next.Next
//
//}
