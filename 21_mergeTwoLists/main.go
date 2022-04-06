package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var l1 ListNode
	l1.Val = 0
	nTmp1 := &l1
	for i := 0; i < 3; i++ {
		var node ListNode
		node.Val = i
		nTmp1.Next = &node
		nTmp1 = &node
	}

	var l2 ListNode
	l2.Val = 0
	nTmp2 := &l2
	for i := 0; i < 2; i++ {
		var node ListNode
		node.Val = i * 2
		nTmp2.Next = &node
		nTmp2 = &node
	}

	result := mergeTwoLists(&l1, &l2)
	fmt.Println(result)
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var list ListNode
	tmp := &list
	tmp1 := list1
	tmp2 := list2
	for {
		if tmp1 == nil && tmp2 == nil {
			break
		}

		var node *ListNode
		if tmp1 != nil && tmp2 != nil {
			if tmp1.Val < tmp2.Val {
				node = tmp1
				// 链表1指向下一个节点
				tmp1 = tmp1.Next
			} else {
				node = tmp2
				// 链表2指向下一个节点
				tmp2 = tmp2.Next
			}
		} else if tmp1 != nil {
			node = tmp1
			// 链表1指向下一个节点
			tmp1 = tmp1.Next
		} else if tmp2 != nil {
			node = tmp2
			// 链表2指向下一个节点
			tmp2 = tmp2.Next
		}

		// 保存值较小的节点
		tmp.Next = node
		// 指向下一个节点
		tmp = node
	}

	return list.Next
}
