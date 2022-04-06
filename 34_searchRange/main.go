package main

import "fmt"

func main() {

	var nums []int
	var n int
	var result []int

	nums = []int{5, 7, 7, 8, 8, 10}
	n = 8

	result = searchRange(nums, n)
	fmt.Println(result)

	nums = []int{2, 2}
	n = 8

	result = searchRange(nums, n)
	fmt.Println(result)

	nums = []int{}
	n = 8

	result = searchRange(nums, n)
	fmt.Println(result)

	nums = []int{2, 2}
	n = 2

	result = searchRange(nums, n)
	fmt.Println(result)
}

//// 在排序数组中查找元素的第一个和最后一个位置
//func searchRange(nums []int, target int) []int {
//	min, max := -1, -1
//
//	for i, v := range nums {
//		// 目标值的元素
//		if v == target {
//			// 第一个目标值的元素
//			if min == -1 {
//				min = i
//			}
//			max = i
//		}
//
//		// 最后一个目标值后面的元素，之后的元素不需要再比较
//		if max != -1 && v != target {
//			break
//		}
//	}
//
//	return []int{min, max}
//}

// 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	min, max := -1, -1
	n := BinarySearchRecursive(nums, target, 0, len(nums)-1)

	if n >= 0 {
		// 二分查找位置向前
		for i := n; i >= 0; i-- {
			if nums[i] == target {
				min = i
			}
			if nums[i] != target {
				break
			}
		}

		// 二分查找位置向后
		for i := n; i < len(nums); i++ {
			if nums[i] == target {
				max = i
			}
			if nums[i] != target {
				break
			}
		}
	}

	return []int{min, max}
}

// 递归
func BinarySearchRecursive(slice []int, value int, low int, high int) int {
	if low > high {
		return -1
	}
	// 中间位置等于 低位位置 + （高位位置-低位位置）/2
	mid := low + (high-low)/2
	// 找到目标值，返回位置
	if slice[mid] == value {
		return mid
	}
	// 当前值小于目标值，则低位位置 改为 中间位置 + 1
	if slice[mid] < value {
		low = mid + 1
	}
	// 当前值大于目标值，则高位位置 改为 中间位置 - 1
	if slice[mid] > value {
		high = mid - 1
	}
	return BinarySearchRecursive(slice, value, low, high)
}
