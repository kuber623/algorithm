package binary_search

// SearchLeftBound 查找升序数组中第一个大于等于目标元素的索引
func SearchLeftBound(nums []int, x int) int {
	if len(nums) == 0 {
		return 0
	}
	return recursionSLB(nums, 0, len(nums)-1, x)
}

func recursionSLB(nums []int, low, high, x int) int {
	if low > high {
		return low
	}

	mid := low + (high-low)/2
	switch {
	case nums[mid] >= x:
		// 如果目标在左半区间（包括 mid 本身），则收缩右边界
		return recursionSLB(nums, low, mid, x)
	case nums[mid] < x:
		// 如果目标在右半区间（排除 mid 本身），则收缩左边界
		return recursionSLB(nums, mid+1, high, x)
	default:
		panic("unreachable!!!")
	}

	return 0
}
