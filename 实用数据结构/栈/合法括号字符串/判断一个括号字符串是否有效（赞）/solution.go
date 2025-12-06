package leetcode_2116

// https://leetcode.cn/problems/check-if-a-parentheses-string-can-be-valid/
// 难度：中等
// 题解：
// 我的题解：贪心 + 左平衡校验 + 右平衡校验
// 灵神题解：https://leetcode.cn/problems/check-if-a-parentheses-string-can-be-valid/solutions/1178043/zheng-fan-liang-ci-bian-li-by-endlessche-z8ac/

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if locked[i] == '0' {
			b[i] = '-'
		}
	}

	balanced := 0
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '(', '-':
			balanced++
		case ')':
			balanced--
			if balanced < 0 {
				return false
			}
		}
	}

	balanced = 0
	for i := len(b) - 1; i >= 0; i-- {
		switch b[i] {
		case ')', '-':
			balanced++
		case '(':
			balanced--
			if balanced < 0 {
				return false
			}
		}
	}

	return true
}
