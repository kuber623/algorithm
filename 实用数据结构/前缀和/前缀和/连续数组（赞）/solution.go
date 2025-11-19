package leetcode_525

// https://leetcode.cn/problems/contiguous-array/
// 难度：中等
// 题解：把数组中的 0 看成 −1，计算和为 0 的最长子数组。

func findMaxLength(nums []int) int {
	ans := 0

	m := make(map[int]int, len(nums)) // 记录前缀和首次出现的索引
	m[0] = -1                         // 神来之笔

	sum := 0
	for i, num := range nums {
		if num == 0 {
			num = -1
		}
		sum += num

		if v, ok := m[sum]; !ok {
			m[sum] = i
		} else {
			ans = max(ans, i-v)
		}
	}

	return ans
}
