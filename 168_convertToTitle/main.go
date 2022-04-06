package main

import "fmt"

func main() {

	var columnNumber int
	//columnNumber = 2147483647
	columnNumber = 78

	result := convertToTitle(columnNumber)
	fmt.Println(result)
}

// Excel表列名称
func convertToTitle(columnNumber int) string {
	// 26个字母
	const cntLetter = 26

	// 结果
	var result string
	// 用于计算字母编码的切片
	var ch []int

	idx := columnNumber
	for idx > 0 {
		// 求余数
		tail := idx % cntLetter

		if tail == 0 {
			// 整除无余数，则用 26 来计算编码
			ch = append(ch, cntLetter)
			// 先减去26，再取整数部分进行下一次循环
			idx = (idx - cntLetter) / cntLetter
		} else {
			// 余数 用来计算编码
			ch = append(ch, tail)
			// 取整数部分进行下一次循环
			idx = idx / cntLetter
		}
	}

	// 循环切片，通过ASCII码计算出对应的字母后进行连接
	for _, v := range ch {
		result = string(v+65-1) + result
	}

	return result
}
