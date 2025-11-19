package leetcode_1685

// https://leetcode.cn/problems/sum-of-absolute-differences-in-a-sorted-array/
// 难度：中等

func getSumAbsoluteDifferences(nums []int) (result []int) {
	n := len(nums)
	result = make([]int, n)

	// 计算前缀和
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + nums[i]
	}

	for i := 0; i < n; i++ {
		result[i] += sum[n] - sum[i] - (n-i)*nums[i]
		result[i] += i*nums[i] - sum[i]
	}

	return
}
