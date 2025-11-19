package leetcode_3511

// https://leetcode.cn/problems/make-a-positive-array/
// 难度：中等

func makeArrayPositive(nums []int) (ans int) {
	// 计算前缀和
	// prefsum[i] 表示数组 nums 在区间 [0, i-1] 的累加和
	prefsum := make([]int, len(nums)+1)
	for i := range nums {
		prefsum[i+1] = prefsum[i] + nums[i]
	}

	// 计算前缀和的后缀最小值
	// suffmin[i] 表示数组 prefsum 在区间 [i, ♾️) 的最小值
	suffmin := make([]int, len(prefsum))
	suffmin[len(suffmin)-1] = prefsum[len(prefsum)-1]
	for i := len(suffmin) - 2; i >= 0; i-- {
		suffmin[i] = min(suffmin[i+1], prefsum[i])
	}

	for i := len(prefsum) - 1; i >= 3; i-- {
		if suffmin[i] < prefsum[i-3] {
			ans++
		}
	}

	return
}
