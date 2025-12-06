package leetcode_868

// https://leetcode.cn/problems/binary-gap/
// 难度：简单

func binaryGap(n int) int {
	gap := 0
	for i, j := 0, -1; n > 0; i++ {
		if n&1 == 1 {
			if j != -1 {
				gap = max(gap, i-j)
			}
			j = i
		}
		n >>= 1
	}
	return gap
}
