package main

import "fmt"

func main() {
	//pushed := []int{1, 2, 3, 4, 5}
	//popped := []int{4, 5, 3, 2, 1}
	//pushed := []int{2, 1, 0}
	//popped := []int{1, 2, 0}
	pushed := []int{4, 0, 1, 2, 3}
	popped := []int{4, 2, 3, 0, 1}
	result := validateStackSequences(pushed, popped)
	fmt.Println(result)
}

// 946. 验证栈序列
func validateStackSequences(pushed []int, popped []int) bool {
	// 栈
	var stack []int

	for {
		// 入栈出栈标记
		isPushed := false
		isPopped := false

		// 栈非空，且栈中最顶端的元素 和 弹出序列的第一个值相同，则出栈并从弹出序列中删除该值
		if len(stack) > 0 && stack[len(stack)-1] == popped[0] {
			// 弹出栈
			stack = stack[:len(stack)-1]
			// 弹出序列非空时，删除该值
			if len(popped) > 0 {
				popped = popped[1:]
				// 出栈标记设为 true
				isPopped = true
			}
		} else if len(pushed) > 0 {
			// 如果推入序列非空，则将推入序列的第一个值压入栈，并从推入序列中删除该值
			stack = append(stack, pushed[0])
			pushed = pushed[1:]
			// 入栈标记设为 true
			isPushed = true
		}

		// 当次循环同时无出栈和入栈处理时，跳出循环
		if !isPushed && !isPopped {
			break
		}
	}

	// 返回栈长度是否为0
	return len(stack) == 0
}
