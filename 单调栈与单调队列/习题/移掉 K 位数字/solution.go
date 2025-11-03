package leetcode_402

// https://leetcode.cn/problems/remove-k-digits
// 难度：中等
// 题解：
// 为了使删除 K 位数字后足够小，需要保证靠前的数字尽可能小。
// 我们可以通过单调栈维护

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	stack := make([]int, 0, len(num))
	for _, v := range num {
		for len(stack) > 0 && top(stack) > int(v-'0') && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, int(v-'0'))
	}

	for k > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	digits := make([]rune, 0, len(stack))
	var i int
	for i = 0; i < len(stack); i++ {
		if stack[i] != 0 {
			break
		}
	}

	if i == len(stack) {
		return "0"
	} else {
		for ; i < len(stack); i++ {
			digits = append(digits, rune(stack[i]-'0'))
		}
	}

	return string(digits)
}

func top(stack []int) int {
	return stack[len(stack)-1]
}
