package main

import "fmt"

func main() {
	nums := []int{1}
	//nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	nums = []int{1, 1, 2, 2}
	result := removeDuplicates(nums)
	fmt.Println(result)
	fmt.Println(nums)
}

//// 移除元素
//func removeDuplicates(nums []int) int {
//	ln := len(nums)
//	for i := 0; i < ln; i++ {
//		for j := i + 1; j < ln+1; j++ {
//			if j == ln {
//				break
//			}
//
//			if nums[j] == nums[i] {
//				if j < ln-1 {
//					for k := j; k < ln-1; k++ {
//						nums[k] = nums[k+1]
//					}
//					ln = ln - 1
//					j = j - 1
//				} else {
//					nums = nums[:i+1]
//					return len(nums)
//				}
//			}
//		}
//	}
//
//	return ln
//}

// 移除元素
func removeDuplicates(nums []int) int {
	ln := len(nums)
	for i := 0; i < ln; i++ {
		val := nums[i]
		left := nums[:i+1]
		for j := i + 1; j < ln+1; j++ {
			if j == ln {
				break
			}

			if nums[j] == val {
				if j < ln-1 {
					nums = append(left, nums[j+1:]...)
					ln = ln - 1
					j = j - 1
				} else {
					nums = left
					return len(nums)
				}
			}
		}
	}

	return len(nums)
}
