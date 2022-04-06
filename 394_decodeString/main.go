package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//s := "xyz12[abc]3[cd]ef"
	s := "3[a2[c]]"
	//s := "3[a]2[bc]"
	result := decodeString(s)
	fmt.Println(result)
}

func decodeString(s string) string {

	// 结果栈
	var resultStack []string
	// 数字栈
	var numStack []string
	// 中括号内字母栈
	var partStack []string

	//var afterLetter bool

	// 中括号内标志
	tmp := ""
	for _, v := range s {
		if unicode.IsLetter(v) || v == '[' {
			// 中括号内标志为 true ，则将字母推入到 中括号内字母栈
			partStack = append(partStack, string(v))
		}
		// 数字推入到 数字栈
		if unicode.IsNumber(v) || v == '[' {
			numStack = append(numStack, string(v))
		}

		// 遇到右中括号，出栈
		repeat := ""
		if v == ']' {
			// 计算 数字栈 中的数值
			num := 0
			// 数字位
			powN := 0
			numStack = numStack[0 : len(numStack)-1]
			for n := len(numStack) - 1; n >= 0; n-- {

				// 各数字位的值进行加法运算
				val, _ := strconv.Atoi(numStack[n])
				num = num + int(math.Pow10(powN))*val
				powN = powN + 1

				if numStack[n] == "[" {
					numStack = numStack[0 : n+1]
					break
				}
				if n == 0 {
					numStack = numStack[0:n]
					break
				}
			}

			// 中括号内字母栈 的字母连接成字符串
			sPart := ""
			for n := len(partStack) - 1; n >= 0; n-- {
				if partStack[n] == "[" {
					partStack = partStack[0:n]
					break
				}

				sPart = partStack[n] + sPart
			}

			if len(partStack) == 0 {
				// 生成 数字栈 中结果 个 字符串
				repeat = strings.Repeat(sPart, num)
			} else {
				tmp = sPart + tmp
				// 生成 数字栈 中结果 个 字符串
				repeat = strings.Repeat(tmp, num)
			}
			resultStack = append(resultStack, repeat)
		}
	}

	// 结果栈中的字母连接成字符串后返回
	return tmp
}

//func decodeString(s string) string {
//
//	// 结果栈
//	var resultStack []string
//	// 数字栈
//	var numStack []int
//	//
//	var partStack []string
//
//	isInBracket := false
//	for _, v := range s {
//		if string(v) == "[" {
//			isInBracket = true
//		} else if isInBracket && string(v) >= "a" && string(v) <= "z" {
//			partStack = append(partStack, string(v))
//		} else if !isInBracket && string(v) >= "a" && string(v) <= "z" {
//			resultStack = append(resultStack, string(v))
//		}
//		if string(v) >= "0" && string(v) <= "9" {
//			val, _ := strconv.Atoi(string(v))
//			numStack = append(numStack, val)
//		}
//		fmt.Println(numStack)
//
//		if string(v) == "]" {
//			num := 0
//			powN := len(numStack) - 1
//			for n := 0; n < len(numStack); n++ {
//				num = num + int(math.Pow10(powN))*numStack[n]
//				powN = powN - 1
//			}
//			fmt.Println(num)
//
//			s := strings.Join(partStack, "")
//			fmt.Println(s)
//
//			repeat := strings.Repeat(s, num)
//			resultStack = append(resultStack, repeat)
//
//			partStack = partStack[0:0]
//			numStack = numStack[0:0]
//			isInBracket = false
//		}
//	}
//
//	return strings.Join(resultStack, "")
//}
