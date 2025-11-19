package binary_search

// BinarySearchLowerBound 返回最小的满足 nums[i] >= target 的 i
// 如果不存在，则返回数组长度
func BinarySearchLowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 区间 [left, right]
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

// BinarySearchUpperBound 返回最大满足 nums[i] < target 的 i
// 如果不存在，则返回
