package main

import "fmt"

//
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var l1 ListNode
	l1.Val = 0
	tmp := &l1

	for i := 0; i < 5; i++ {
		var node ListNode
		node.Val = i + 1
		tmp.Next = &node
		tmp = tmp.Next
	}

	result := middleNode(&l1)
	fmt.Println(5 / 2)
	fmt.Println(result)
}

//// 链表的中间结点
//// 两次循环
//func middleNode(head *ListNode) *ListNode {
//	// 获取链表的结点总数
//	node := head
//	length := 1
//	for ; ; length = length + 1 {
//		if node.Next == nil {
//			break
//		}
//
//		node = node.Next
//	}
//
//	// 中间结点位置
//	iTarget := length / 2
//
//	// 循环找到中间结点
//	var result *ListNode
//	node = head
//	for i := 0; ; i++ {
//		if i == iTarget {
//			result = node
//			break
//		}
//
//		node = node.Next
//	}
//
//	return result
//}

// 链表的中间结点
// 单次循环
func middleNode(head *ListNode) *ListNode {
	// 结点指针切片，存放各结点指针
	var slice []*ListNode

	// 获取链表的结点总数并将各结点的指针保存到切片中。
	node := head
	length := 1
	for ; ; length = length + 1 {
		slice = append(slice, node)

		if node.Next == nil {
			break
		}

		node = node.Next
	}

	// 目标结点位置
	iTarget := length / 2

	// 从结点切片中找到对应的结点并返回
	return slice[iTarget]
}
