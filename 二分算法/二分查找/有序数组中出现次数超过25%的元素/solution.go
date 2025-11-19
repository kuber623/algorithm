package leetcode_1287

// https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array
// 难度：简单
// 题解：
// 看艾神题解：https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/solutions/3067559/olog-n-er-fen-cha-zhao-zheng-que-xing-zh-5mu9/

func findSpecialInteger(arr []int) int {
	left, right := 0, len(arr)-1
	mid := left + (right-left)>>1
	if countInteger(arr, arr[mid]) >= len(arr)/4 {
		return mid
	}

	lmid := left + (mid-1-left)>>1
	if countInteger(arr[:mid], arr[lmid]) >= len(arr)/4 {
		return lmid
	}

	rmid := mid + 1 + (right-mid-1)>>1
	if countInteger(arr[mid+1:], arr[rmid]) >= len(arr)/4 {
		return rmid
	}

	return -1
}

// countInteger 计算整数 target 在 nums 中出现的个数
func countInteger(nums []int, target int) int {
	start := searchLowerBound(nums, target)
	if start == len(nums) || nums[start] != target {
		return 0
	}
	end := searchLowerBound(nums, target+1)
	return end - start
}

func searchLowerBound(nums []int, target int) int {
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
