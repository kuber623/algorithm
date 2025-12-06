package leetcode_1342

import "math/bits"

// https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-to-zero/
// 难度：简单

func numberOfSteps(num int) int {
	return max(bits.Len(uint(num))+bits.OnesCount(uint(num))-1, 0)
}
