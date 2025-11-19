package leetcode_1177

import "math/bits"

// https://leetcode.cn/problems/can-make-palindrome-from-substring/description/
// 难度：中等
// 题解：异或和
// 通过异或计算每个字母的奇偶性，由于偶数个数的字母可以通过重排序到字母串两端构成回文，故只需要关注奇数个数的字母
// 通过异或和可以快速计算任意子串中字母的奇偶性，得到奇数字母的个数后
// 假设有 m 种字母出现奇数次，只需要修改 m/2 个字母使子串变成回文串

func canMakePaliQueries(s string, queries [][]int) []bool {
	ans := make([]bool, 0, len(queries))

	n := len(s)
	xor := make([]uint32, n+1) // 异或和
	for i, c := range s {
		xor[i+1] = xor[i] ^ (1 << (c - 'a'))
	}

	for _, query := range queries {
		l, r, k := query[0], query[1], query[2]
		x := bits.OnesCount32(xor[r+1] ^ xor[l])
		ans = append(ans, x/2 <= k)
	}

	return ans
}
