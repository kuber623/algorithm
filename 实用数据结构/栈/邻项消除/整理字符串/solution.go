package leetcode_1544

// https://leetcode.cn/problems/make-the-string-great/
// 难度：简单
// 题解：栈

func makeGood(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1]-s[i] == 32 || s[i]-stack[len(stack)-1] == 32 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
