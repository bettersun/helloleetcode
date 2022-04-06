package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "abcde"
	goal := "cdeab"

	result := rotateString(s, goal)
	fmt.Println(result)

}

// 796. 旋转字符串
func rotateString(s string, goal string) bool {

	return strings.Index(s+s, goal) >= 0
}
