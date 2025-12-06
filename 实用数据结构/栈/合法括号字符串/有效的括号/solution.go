package leetcode_20

// https://leetcode.cn/problems/valid-parentheses/
// 难度：简单

var brackets = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		bracket, ok := brackets[c]
		if !ok {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != bracket {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
