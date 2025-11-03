package binary_search

// Search 查找目标值在升序数组中的索引，不存在返回 -1
func Search(nums []int, x int) int {
	return recursionBS(nums, 0, len(nums)-1, x)
}

func recursionBS(nums []int, low, high, x int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2
	switch {
	case nums[mid] == x:
		return mid
	case nums[mid] > x:
		return recursionBS(nums, low, mid-1, x)
	case nums[mid] < x:
		return recursionBS(nums, mid+1, high, x)
	default:
		panic("unreachable!!!")
	}
}
