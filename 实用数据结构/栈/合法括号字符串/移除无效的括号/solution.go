package leetcode_1249

import "strings"

// https://leetcode.cn/problems/minimum-remove-to-make-valid-parentheses/
// 难度：中等

func minRemoveToMakeValid(s string) string {
	ans := make([]rune, 0, len(s))

	cntL, cntR := 0, strings.Count(s, ")") // 表示未匹配的左右括号数量
	for _, c := range s {
		switch c {
		case '(':
			if cntL >= cntR {
				continue
			}
			cntL++
		case ')':
			cntR--
			if cntL <= 0 {
				continue
			}
			cntL--
		}
		ans = append(ans, c)
	}

	return string(ans)
}
