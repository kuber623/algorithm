package leetcode_316

// https://leetcode.cn/problems/remove-duplicate-letters
// 难度：中等
// 题解：

func removeDuplicateLetters(s string) string {
	// 记录每个字母最后出现的位置
	m := make(map[rune]int, 26)
	for i, v := range s {
		if i == 0 {
			m[v] = i
			continue
		}
		if rune(s[i-1]) != v {
			m[v] = i
		}
	}

	stack := make([]int, 0, 26) // 单调递增栈
	exist := make(map[rune]bool, 26)
	for i, v := range s {
		// 如果当前字母已存在于单调栈中，直接跳过
		if exist[v] {
			continue
		}

		for len(stack) > 0 {
			r := topRune(s, stack)
			// 将所有比当前字母大且后面尚存在的字母出栈
			if r > v && top(stack) < m[r] && i < m[r] {
				stack = pop(stack)
				exist[r] = false
				continue
			}
			break
		}

		// 将当前字母压栈
		stack = append(stack, i)
		exist[v] = true
	}

	res := make([]byte, len(stack))
	for i, v := range stack {
		res[i] = s[v]
	}

	return string(res)
}

func top(stack []int) int {
	return stack[len(stack)-1]
}

func pop(stack []int) []int {
	return stack[:len(stack)-1]
}

func topRune(s string, stack []int) rune {
	return rune(s[top(stack)])
}
