package main

import (
	"fmt"
)

func main() {
	x := -2345
	//x = 1534236469
	result := reverse(x)
	fmt.Println(result)
}

// 整数反转
func reverse(x int) int {
	// 创建一个切片，存放 [1, 10, ..., 10 的 目标数值最大位数 次方]
	var div []int
	div = append(div, 1)
	for i := 10; ; i = i * 10 {
		// 超出最大位数，跳出循环
		if x/i == 0 {
			break
		}
		// 未超出最大位数，将 10 的位数次方添加到切片
		div = append(div, i)
	}

	sum := 0
	for i, j := len(div)-1, 0; i >= 0; i, j = i-1, j+1 {
		// 从高位获取各位的数字
		tmp := x / div[i] % 10
		// 高位的数字 乘以(反转后)对应的低位的 10 的 j 次方后，累加到结果
		// 如果环境不允许存储 64 位整数（有符号或无符号），此处应该会发生异常（未验证）。
		sum = sum + div[j]*tmp

		// 右移31位，如果大于0，则是正数越界；如果小于-1，则是负数越界。
		if sum>>31 > 0 || sum>>31 < -1 {
			return 0
		}
	}

	return sum
}

//func reverse(x int) int {
//	// 创建一个切片，存放 [1, 10, ..., 10 的 目标数值最大位数 次方]
//	var div []int
//	div = append(div, 1)
//	for i := 10; ; i = i * 10 {
//		if x/i == 0 {
//			break
//		}
//
//		div = append(div, i)
//	}
//
//	var dec []int
//	for i := len(div) - 1; i >= 0; i-- {
//		tmp := x / div[i] % 10
//		dec = append(dec, tmp)
//	}
//
//	sum := 0
//	for i := 0; i < len(div); i++ {
//		tmp := div[i] * dec[i]
//		sum = sum + tmp
//
//		if sum>>31 > 0 || sum>>31 < -1 {
//			return 0
//		}
//	}
//
//	return sum
//}
