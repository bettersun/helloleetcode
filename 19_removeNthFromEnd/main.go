package main

import "fmt"

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

	var result *ListNode
	//result = removeNthFromEnd(&l1, 1)
	//fmt.Printf("%+v", result)
	//result = removeNthFromEnd(&l1, 2)
	//fmt.Printf("%+v", result)
	//result = removeNthFromEnd(&l1, 3)
	//fmt.Printf("%+v", result)
	result = removeNthFromEnd(&l1, 4)
	fmt.Printf("%+v", result)
}

//// 删除链表的倒数第 N 个结点
//// 单次循环
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	// 节点指针切片，存放各节点
//	var slice []*ListNode
//
//	// 获取链表节点数并将各节点的指针保存到切片中。
//	node := head
//	length := 1
//	for ; node.Next != nil; length = length + 1 {
//		slice = append(slice, node)
//		node = node.Next
//	}
//
//	// 目标节点位置
//	iTarget := length - n
//
//	if iTarget == 0 {
//		// 目标节点为头节点，则指向头节点的下一个节点
//		head = head.Next
//	} else if iTarget == length-1 {
//		// 目标节点为最后一个节点，则目标节点的前一个节点的下一个节点设置为 nil
//		slice[iTarget-1].Next = nil
//	} else {
//		// 目标节点的前一个节点的下一个节点 设置为 目标节点的下一个节点
//		slice[iTarget-1].Next = slice[iTarget].Next
//	}
//
//	return head
//}

// 删除链表的倒数第 N 个结点
// 两次循环
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head

	// 获取链表节点数
	length := 1
	for ; node.Next != nil; length = length + 1 {
		node = node.Next
	}

	// 目标节点位置
	iTarget := length - n

	// 前一个节点
	var prev *ListNode
	// 当前节点
	current := head
	// 循环到 目标节点位置 为止
	for i := 0; i <= iTarget; i++ {
		if iTarget == 0 {
			// 目标节点为头节点，则指向头节点的下一个节点
			head = current.Next
		} else if i == iTarget {
			// 已是目标节点，则将前一个节点的下一个节点设置为当前节点的下一个节点
			prev.Next = current.Next
		}

		// 当前节点赋给前一个节点（用于下次循环）
		prev = current
		// 下一个节点再作为当前节点（用于下次循环）
		current = current.Next
	}

	return head
}
