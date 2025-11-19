package leetcode_3583

// https://leetcode.cn/problems/count-special-triplets/
// 难度：中等
// 题解：枚举中间解法 or 一次遍历解法

// 枚举中间解法
// 预处理 pref 和 suff 分别维护左侧和右侧的数字计数，然后枚举中间值 num，查询 2 * num 在 pref 和 suff 中的计数
// 两者间的乘积为以 num 为中间值的三元组个数，即 sum += pref[2*num] + suff[2*num]
// 遍历过程中需要动态维护 pref 和 suff
func specialTripletsEnumMid(nums []int) (ans int) {
	const mod int = 1e9 + 7

	n := len(nums)
	pref, suff := make(map[int]int), make(map[int]int)
	pref[nums[0]]++
	for i := n - 1; i > 0; i-- {
		suff[nums[i]]++
	}

	// 枚举中间，维护两边
	for i := 1; i < n-1; i++ {
		num := nums[i]
		suff[num]--
		ans += pref[2*num] * suff[2*num]
		pref[num]++
	}

	return ans % mod
}

// 一次遍历解法
func specialTripletsOneTraversal(nums []int) (ans int) {
	const mod int = 1e9 + 7

	m1, m2 := make(map[int]int), make(map[int]int) // m1 记录 nums[i] 的
	for _, num := range nums {
		// 把当前数字视为 nums[k]
		if num%2 == 0 {
			ans += m2[num/2]
		}
		// 把当前数字视为 nums[j]
		m2[num] += m1[num*2]
		// 把当前数字视为 nums[i]
		m1[num]++
	}

	return ans % mod
}
