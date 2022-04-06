package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 5, 3, 3, 4, 3, 2, 4, 2}
	//nums := []int{1, 2, 10, 6, 8, 9, 12, 200, 42, 23}
	result := containsDuplicate(nums)
	fmt.Println(result)
}

// 存在重复元素
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

//// 存在重复元素
//func containsDuplicate(nums []int) bool {
//	m := make(map[int]int)
//	for _, v := range nums {
//		if _, ok := m[v]; ok {
//			return true
//		}
//
//		m[v] = 0
//	}
//	return false
//}
//
//// 存在重复元素
//func containsDuplicate(nums []int) bool {
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//
//	return false
//}
