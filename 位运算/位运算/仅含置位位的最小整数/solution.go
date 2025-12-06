package leetcode_3370

import (
	"math"
)

// https://leetcode.cn/problems/smallest-number-with-all-set-bits/
// 难度：简单

func smallestNumber(n int) int {
	return 1<<int(math.Log2(float64(n))+1) - 1
}
