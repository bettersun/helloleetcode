package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4, 5, 6}
	result := findMedianSortedArrays(nums1, nums2)

	fmt.Println(result)
}

// 寻找两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 数组元素个数
	ln1 := len(nums1)
	ln2 := len(nums2)
	ln := ln1 + ln2

	// 中间下标
	mid := 0
	isEven := false // 数组元素总个数是否为偶数标志
	if ln%2 == 1 {
		mid = (ln - 1) / 2
	} else {
		mid = ln / 2
		isEven = true
	}

	m := 0
	n := 0
	current := 0
	prev := 0

	for i := 0; i <= mid; i++ {
		// 上一个值
		prev = current

		if m < ln1 && n < ln2 {
			// 比较后小的值设置给当前值
			if nums1[m] < nums2[n] {
				current = nums1[m]
				m = m + 1
			} else {
				current = nums2[n]
				n = n + 1
			}
		} else if m < ln1 {
			// 当前值
			current = nums1[m]
			m = m + 1
		} else if n < ln2 {
			// 当前值
			current = nums2[n]
			n = n + 1
		}
	}

	if isEven {
		// 元素总个数为偶数时，取当前值和上一个值的和的平均值
		return float64(current+prev) / 2
	} else {
		// 元素总个数为奇数时，取当前值
		return float64(current)
	}

	return 0
}
