package main

import (
	"fmt"
)

func main() {

	s := "AB"
	result := titleToNumber(s)
	fmt.Println(result)
}

// Excel 表列序号
func titleToNumber(columnTitle string) int {
	const cntLetter = 26
	var result int

	n := 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		// 字母转换成对应的 26 进制的数字
		// A : 1 , B:2 , Z:26
		val := (int(columnTitle[i]) - 65 + 1) * n
		// 各位结果相加
		result = result + val

		// 从低位到高位，是 26 的 n 次方
		// 最右端：26 的 0 次方
		// 右边第二位：26 的 1 次方
		// 右边第三位：26 的 2 次方
		n = n * cntLetter
	}

	return result
}
