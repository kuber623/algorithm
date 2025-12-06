package leetcode_476

import "math/bits"

// https://leetcode.cn/problems/number-complement/
// 难度：简单

func findComplement(num int) int {
	return num ^ (1<<(bits.Len32(uint32(num))) - 1)
}
