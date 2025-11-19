package leetcode_2364

// https://leetcode.cn/problems/count-number-of-bad-pairs/
// 难度：中等
// 题解：
// 根据坏数对的定义：j - i != nums[j] - nums[i] 通过移项可以得到：nums[i] - i != nums[j] - j
// 我们可以将好数对定义为 nums[i] - i = nums[j] - j，即对数组中每个元素减去索引号后，如果两个数相等，则这两个数可以组成好数对
// 在遍历数组时，我们可以记录值为 nums[i] - i 的个数，并累加到好数对总数当中，最终得到总的好数对个数
// 而坏数对数量可以通过「总数对数量 - 好数对数量」得到

func countBadPairs(nums []int) int64 {
	n := len(nums)

	// 计算好数对数量
	good := int64(0)
	cnt := make(map[int]int64)
	for i := 0; i < n; i++ {
		good += cnt[nums[i]-i]
		cnt[nums[i]-i]++
	}

	// 计算坏数对数量
	bad := int64(n*(n-1)/2) - good

	return bad
}
