package main

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
}

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for i, v := range nums {
		if v == 0 {
			tmp := nums[i]
			nums[i] = nums[left]
			nums[left] = tmp
			left = left + 1
		}

		if v == 2 {
			tmp := nums[i]
			nums[i] = nums[right]
			nums[right] = tmp
			right = right - 1
		}

		if i >= right {
			break
		}
	}
}
