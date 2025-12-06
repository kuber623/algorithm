package leetcode_32

// https://leetcode.cn/problems/longest-valid-parentheses/
// 难度：困难

func longestValidParentheses(s string) (ans int) {
	stack := make([]int, 0, len(s))

	for i, c := range s {
		if c == ')' && len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, i)
	}

	if len(stack) == 0 {
		return len(s)
	}

	for i := 1; i < len(stack); i++ {
		ans = max(ans, stack[i]-stack[i-1]-1)
	}
	ans = max(ans, stack[0], len(s)-stack[len(stack)-1]-1)

	return
}
