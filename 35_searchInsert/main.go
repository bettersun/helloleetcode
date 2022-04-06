package main

import "fmt"

func main() {

	nums := []int{6}
	target := 4

	n := SearchInsert(nums, target)
	fmt.Println(n)
	fmt.Println(nums)
}

// 搜索插入位置
func SearchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	n := -1
	for low <= high {
		// 中间位置等于 低位位置 + （高位位置-低位位置）/2
		mid := low + (high-low)/2

		// 找到目标值，记录位置，跳出循环
		if target == nums[mid] {
			n = mid
			break
		}

		// 当前值小于目标值，则低位位置 改为 中间位置 + 1
		if target < nums[mid] {
			high = mid - 1
		}
		// 当前值大于目标值，则高位位置 改为 中间位置 - 1
		if target > nums[mid] {
			low = mid + 1
		}
	}

	// 未找到目标值，则应在上面循环结束的位置添加目标值
	if n == -1 {
		n = low
	}

	return n
}
