package leetcode_34

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 难度：中等
// 题解：

func searchRange(nums []int, target int) []int {
	start := doSearchLowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := doSearchLowerBound(nums, target+1)
	return []int{start, end - 1}
}

func doSearchLowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
