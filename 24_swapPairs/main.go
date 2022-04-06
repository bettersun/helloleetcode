package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 ListNode
	//l1.Val = 12
	nTmp1 := &l1
	for i := 0; i < 7; i++ {
		var node ListNode
		node.Val = i + 1
		nTmp1.Next = &node

		fmt.Println(nTmp1)
		nTmp1 = &node
	}

	result := swapPairs(&l1)
	fmt.Println(result)

	nTmp2 := result
	for i := 0; i < 7; i++ {
		fmt.Println(nTmp2)
		nTmp2 = nTmp2.Next
	}
}

// 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	// 定义一个新的结点链表，用于把整理好的结点挂在上面。
	var newHead ListNode
	// 新的结点链表的指针，用于记录 整理好的结点的位置
	ptr := &newHead

	// 目标链表的指针
	node := head
	for i := 0; node != nil; i++ {
		// 当前节点为偶数结点位置时
		if i%2 == 0 {
			// 如果当前节点的 Next（下一个结点）不为 空时，交换结点
			if node.Next != nil {
				// 定义临时变量指向当前结节的 Next （下一个结点）
				tmp := node.Next
				// 当前结节的 Next （下一个结点） 指向 当前结节的 Next （下一个结点）的 Next（下一个结点）
				node.Next = node.Next.Next
				// 当前结节的 Next （下一个结点）的 Next （下一个结点）设为当前结点
				tmp.Next = node
				// 挂到新的结点链表上
				ptr.Next = tmp
			} else {
				// 如果当前节点的 Next（下一个结点）不为 空时，
				// 将当前节点的 Next（下一个结点）挂到新的结点链表上
				ptr.Next = node
			}

			// 处理下一个结点
			node = node.Next
		}

		ptr = ptr.Next
	}

	return newHead.Next
}
