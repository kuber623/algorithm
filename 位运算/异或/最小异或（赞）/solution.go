package leetcode_2428

import "math/bits"

// https://leetcode.cn/problems/minimize-xor/
// 难度：中等

func minimizeXor(num1 int, num2 int) int {
	k := bits.OnesCount(uint(num2)) - bits.OnesCount(uint(num1))
	for k > 0 {
		// 最低位的 0 变为 1
		num1 += lowbit(^num1)
		k--
	}
	for k < 0 {
		// 最低位的 1 变为 0
		num1 -= lowbit(num1)
		k++
	}
	return num1
}

func lowbit(x int) int {
	return x & (-x)
}
