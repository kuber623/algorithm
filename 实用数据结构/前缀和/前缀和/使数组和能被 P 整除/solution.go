package leetcode_1590

import "math"

// https://leetcode.cn/problems/make-sum-divisible-by-p/
// 难度：中等
// 题解：
// 本题与 leetcode 974 相反，需要判断是否存在子数组的余与整个数组的余相同
// 如果存在，则返回最短子数组长度；反正返回 -1
// 根据同余定理，两个数之差的余数等于余数之差的余数
// 假设 X Y 问别为两个前缀和对 P 的求余结果，即 (S[j] - S[i]) mod P == K

func minSubarray(nums []int, p int) int {
	// 求数组之和对 P 的余数
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	k := sum % p
	if k == 0 {
		return 0
	}

	ans := math.MaxInt
	sum = 0
	lastMod := map[int]int{
		0: -1,
	}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		x := sum % p         // 当前前缀和余数
		y := (x - k + p) % p // 目标前缀和余数
		if v, ok := lastMod[y]; ok {
			ans = min(ans, i-v)
		}
		lastMod[x] = i
	}

	if ans == math.MaxInt || ans == len(nums) {
		return -1
	}
	return ans
}
