package main

import "fmt"

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// [1,9,9,9]
	var l1 ListNode
	l1.Val = 1
	nTmp1 := &l1
	for i := 0; i < 3; i++ {
		var node ListNode
		node.Val = 9
		nTmp1.Next = &node
		nTmp1 = &node
	}

	// [2,9,9]
	var l2 ListNode
	l2.Val = 2
	nTmp2 := &l2
	for i := 0; i < 2; i++ {
		var node ListNode
		node.Val = 9
		nTmp2.Next = &node
		nTmp2 = &node
	}

	// [3,8,9,0,1]
	// 9991 + 992 = 10983
	result := addTwoNumbers(&l1, &l2)
	fmt.Println(result)
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var node ListNode
	nTmp := &node
	nTmp1 := l1
	nTmp2 := l2

	// 保存加法进位的值
	dec := 0
	for nTmp1 != nil || nTmp2 != nil || dec == 1 {
		var node ListNode

		// 同一层的节点进行加法运算
		tmp := 0
		if nTmp1 != nil && nTmp2 != nil {
			tmp = nTmp1.Val + nTmp2.Val + dec
		} else if nTmp1 != nil {
			tmp = nTmp1.Val + dec
		} else if nTmp2 != nil {
			tmp = nTmp2.Val + dec
		} else {
			tmp = dec
		}

		// 当前节点的值
		node.Val = tmp % 10
		nTmp.Next = &node
		nTmp = &node

		// 进位的值
		dec = tmp / 10
		// 循环条件，最后一位进位时，dec 置为 0
		if nTmp1 == nil && nTmp2 == nil && dec == 1 {
			dec = 0
		}

		// 指向下一个节点
		if nTmp1 != nil {
			nTmp1 = nTmp1.Next
		}
		// 指向下一个节点
		if nTmp2 != nil {
			nTmp2 = nTmp2.Next
		}
	}

	return node.Next
}
