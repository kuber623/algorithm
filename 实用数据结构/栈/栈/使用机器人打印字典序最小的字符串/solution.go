package leetcode_2434

import "math"

// https://leetcode.cn/problems/using-a-robot-to-print-the-lexicographically-smallest-string/
// 难度：中等

func robotWithString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	ans := make([]byte, 0, n)

	// minimum[i] 表示 [i, +♾️) 区间中最小字符
	minimum := make([]byte, n+1)
	minimum[n] = math.MaxUint8
	for i := n - 1; i >= 0; i-- {
		minimum[i] = min(s[i], minimum[i+1])
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < n; i++ {
		stack = append(stack, s[i])
		for len(stack) > 0 && stack[len(stack)-1] <= minimum[i] {
			ans = append(ans, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return string(ans)
}
