package main

import "fmt"

func main() {
	nums := []int{2, 1, 2, 2, 3, 0, 4, 2}
	val := 2

	nums = []int{3}
	val = 3

	//nums = []int{3, 3}
	//val = 3
	//
	//nums = []int{3, 4}
	//val = 3
	//
	//nums = []int{3, 4}
	//val = 4
	fmt.Println(removeElement(nums, val))
}

// 移除元素
func removeElement(nums []int, val int) int {
	ln := len(nums)
	result := ln
	for i := 0; i < ln; i++ {
		if nums[i] == val {

			// 从当前元素的下一个元素开始内层循环
			for j := i + 1; j < ln+1; j++ {
				// 已超出最大下标，截取到外层下标的前一个元素，返回长度
				if j == ln {
					nums = nums[0:i]
					return len(nums)
				}

				// 内层循环的值也等于目标值
				if nums[j] == val {
					if j < ln-1 {
						// 非到最大下标，继续下一次循环
						continue
					} else {
						// 已是最大下标，截取到外层下标的前一个元素，返回长度
						nums = nums[0:i]
						return len(nums)
					}
				}

				// 和不等于目标值的元素交换
				nums[i] = int(nums[j])
				nums[j] = val
				// 结果长度减1
				result = result - 1
				break
			}
		}
	}

	return result
}
