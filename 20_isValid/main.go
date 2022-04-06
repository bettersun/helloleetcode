package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "()[)[]{])[]{{)[]{}"
	//s := "()[]{[]()[]}"
	s := "(([]){})"
	result := isValid(s)
	fmt.Println(result)
}

// 有效的括号
func isValid(s string) bool {
	tmp := s

	// 无限循环
	for {
		// 替换前长度
		ln := len(tmp)

		// 对三种括号进行替换
		tmp = strings.ReplaceAll(tmp, "()", "")
		tmp = strings.ReplaceAll(tmp, "[]", "")
		tmp = strings.ReplaceAll(tmp, "{}", "")

		// 替换前后长度相同时，跳出循环
		if ln == len(tmp) {
			break
		}
	}

	// 替换后结果为空字符串，则返回 true ，否则返回 false 。
	if tmp == "" {
		return true
	}
	return false
}
