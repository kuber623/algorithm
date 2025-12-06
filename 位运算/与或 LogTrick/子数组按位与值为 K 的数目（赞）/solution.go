package leetcode_3209

import "sort"

// https://leetcode.cn/problems/number-of-subarrays-with-and-value-of-k/
// 难度：困难
// 题解：
// 参考灵神 LogTrick 模板 https://www.bilibili.com/video/BV1Qx4y1E7zj/
// 枚举右端点，从右向左遍历左端点，遍历左端点时对枚举数进行 & 操作
// 因此在遍历过程中，每个左端点 nums[L] 维护 nums[L] & nums[L+1] & ... & nums[R] 的值
// 由于需要进行两层遍历，整体的算法复杂度为 O(n^2) 会超时，因此需要进行优化：
// 遍历左端点时，如果当前值在 & 操作后没有发生变化，即可中断本次循环的后续操作，因为后续操作不会使 nums[L] 值出现变化
//
// 最后由于 nums[L] ... nums[R] 是单调非递减数组，因此可以通过二分搜索快速找到数组中值为 K 的数量

func countSubarrays(nums []int, k int) (ans int64) {
	for r, x := range nums {
		for l := r - 1; l >= 0 && nums[l]&x != nums[l]; l-- {
			nums[l] &= x
		}
		arr := nums[:r+1]
		ans += int64(sort.SearchInts(arr, k+1) - sort.SearchInts(arr, k))
	}
	return
}
