package main

import (
	"fmt"
)

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}

// 最后一个单词的长度
func lengthOfLastWord(s string) int {
	//// 去掉末尾空格
	//tmp := strings.TrimSpace(s)
	//// 空格分割成数组
	//slice := strings.Split(tmp, " ")
	//ln := len(slice)
	//// 取最后一个元素的长度
	//return len(slice[ln-1])

	isLetter := false
	c := 0
	for i := len(s) - 1; i >= 0; i-- {
		if !isLetter && s[i] == ' ' {
			continue
		}
		if !isLetter && s[i] != ' ' {
			isLetter = true
		}
		if isLetter && s[i] == ' ' {
			break
		}
		if isLetter && s[i] != ' ' {
			c = c + 1
		}
	}
	return c
}

// 反转字符串
func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
