package leetcode_1915

// https://leetcode.cn/problems/number-of-wonderful-substrings/
// 难度：中等
// 题解：

func wonderfulSubstrings(word string) int64 {
	n := len(word)

	xor := make([]uint16, n+1)
	for i, c := range word {
		xor[i+1] = xor[i] ^ (1 << (c - 'a'))
	}

	ans := 0
	m := [1024]int{}
	for i := 0; i <= n; i++ {
		ans += m[xor[i]]
		for k := 0; k < 10; k++ {
			ans += m[xor[i]^(1<<k)]
		}
		m[xor[i]]++
	}

	return int64(ans)
}
