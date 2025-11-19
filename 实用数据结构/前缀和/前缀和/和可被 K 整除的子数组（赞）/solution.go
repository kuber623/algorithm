package leetcode_974

// https://leetcode.cn/problems/subarray-sums-divisible-by-k/
// 难度：中等
// 题解：
// 通常涉及连续子数组求和问题的时候，可以使用前缀和来解决，子数组和可以视为前后两个前缀和之差
// 假设 S[j] 和 S[i] 分别为 [0, j] 和 [0, i] 区间的前缀和
// 如果希望区间为 [i+1, j] 的子数组和能被 K 整除，即 (S[j] - S[i]) mod K == 0
// 根据「同余定理」 S[j] mod K == S[i] mod K
// 最后通过「右枚举左维护」得到相同余数的前缀和个数，进而得到和能被 K 整除的子数组个数

func subarraysDivByK(nums []int, k int) int {
	m := make(map[int]int, k)
	m[0] = 1

	ans := 0
	sum := 0
	mod := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		mod = (sum%k + k) % k // 确保负数求余后得到正整数
		ans += m[mod]
		m[mod]++
	}

	return ans
}
