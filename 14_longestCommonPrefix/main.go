package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}
	//strs := []string{"cir", "car"}

	result := longestCommonPrefix(strs)
	fmt.Println(result)
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 获取最短字符串的长度
	minlen := 0
	for i, v := range strs {
		ln := len(v)
		if i == 0 {
			minlen = ln
		}
		if ln < minlen {
			minlen = ln
		}
	}

	var result []string
	// 不同标志
	isDiff := false
	// 从0到最短字符串的长度循环
	for m := 0; m < minlen; m++ {
		// 第一个字符串的第 m 个字符
		tmp := strings.Split(strs[0], "")[m]
		c := 0

		// 从第二个字符串循环
		for i := 1; i < len(strs); i++ {
			// 每个字符串的第 m 个文字是否与第一个字符串的第 m 个文字相同
			if strings.Split(strs[i], "")[m] == tmp {
				// 相同则计数加1
				c = c + 1
			} else {
				// 不相同则设置 不同标志 为 true 后，跳出内层循环
				isDiff = true
				break
			}
		}

		// 不同标志 为 true 时，跳出外层循环
		if isDiff {
			break
		}

		// 相同的文字添加到结果切片
		if c == len(strs)-1 {
			result = append(result, tmp)
		}
	}

	// 结果切片中的文字连接成一个字符串，返回
	return strings.Join(result, "")
}
