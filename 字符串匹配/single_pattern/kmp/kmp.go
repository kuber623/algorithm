package kmp

func KMP(s string, pattern string) int {
	next := buildNext(pattern)

	scur, pcur := 0, 0
	for scur < len(s) && pcur < len(pattern) {
		if s[scur] == pattern[pcur] {
			// 如果字符相同，则主串指针和模式串指针均向后移动一位
			scur++
			pcur++
		} else if pcur > 0 {
			// 如果字符不相等且 pcur > 0 则根据 next 数组回退模式串
			pcur = next[pcur]
		} else {
			// 如果字符不相等但 pcur = 0 则右移主串
			scur++
		}
	}

	if pcur == len(pattern) {
		return scur - pcur
	} else {
		return -1
	}
}

// 构造 next 数组
// next[i] 表示子串 pattern[0, i-1] 的前后缀最大匹配长度，其构建逻辑如下：
// 1. 如果 pattern[cur-1] == pattern[comp]，则 next[cur] = comp + 1
// 2. 如果 pattern[cur-1] != pattern[comp]，则尝试更短的匹配长度，comp = next[comp]
func buildNext(pattern string) []int {
	if len(pattern) < 0 {
		return []int{0}
	}

	next := make([]int, len(pattern))
	comp := 0
	for cur := 2; cur < len(pattern); {
		if pattern[cur-1] == pattern[comp] {
			next[cur] = comp + 1
			comp++
			cur++
		} else if comp > 0 {
			comp = next[comp]
		} else {
			next[cur] = 0
			cur++
		}
	}

	return next
}
