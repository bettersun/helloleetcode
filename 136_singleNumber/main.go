package main

import "fmt"

func main() {

	n := []int{4, 1, 2, 1, 2}
	result := singleNumber(n)
	fmt.Println(result)
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for _, v := range nums {
		//result = result ^ v
		result ^= v
	}
	return result
}
