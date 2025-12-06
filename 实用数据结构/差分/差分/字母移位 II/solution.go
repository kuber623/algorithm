package leetcode_2381

// https://leetcode.cn/problems/shifting-letters-ii/
// 难度：中等
// 题解：差分

func shiftingLetters(s string, shifts [][]int) string {
	// 构造差分数组
	diff := make([]int, len(s)+1)
	diff[0] = int(s[0] - 'a')
	for i := 1; i < len(s); i++ {
		diff[i] = int(s[i]) - int(s[i-1])
	}

	// 通过差分数组进行快速区间元素增加和减少
	for _, shift := range shifts {
		begin, end, direction := shift[0], shift[1], shift[2]
		delta := direction
		if direction == 0 {
			delta = -1
		}

		diff[begin] += delta
		diff[end+1] -= delta
	}

	// 还原字符串
	ans := make([]byte, len(s))
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += diff[i]
		ans[i] = byte((sum%26+26)%26 + 'a')
	}

	return string(ans)
}
