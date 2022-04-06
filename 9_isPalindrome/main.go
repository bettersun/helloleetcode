package main

import (
	"fmt"
)

func main() {
	result := isPalindrome(1234321)

	if result {
		fmt.Println("回文数")
	} else {
		fmt.Println("非回文数")
	}
}

// 9. 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var div []int
	for i := 10; ; i = i * 10 {
		if x/i == 0 {
			break
		}
		div = append(div, i)
	}

	ln := len(div)
	for i := 0; i < ln; i++ {
		if i < ln {
			if x%div[i]/(div[i]/10) != x/div[ln-1]%10 {
				return false
			}
		}
		ln = ln - 1
	}
	return true
}
