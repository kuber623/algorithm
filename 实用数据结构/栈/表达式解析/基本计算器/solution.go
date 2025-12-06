package leetcode_224

// https://leetcode.cn/problems/basic-calculator/
// 难度：困难

func calculate(s string) int {
	stack := make([]int, 0) // 存放表达式的结果

	nums := make([]int, 0)
	expr := make([]byte, 0)

	for _, ch := range s {
		switch ch {
		case '+':
			stack = append(stack, number(nums))
		case '-':
		case '*':
		case '/':
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			nums = append(nums, int(ch-'0'))
		case '(':
		case ')':
		}
	}

	return 0
}

func number(nums []int) (ans int) {
	for _, num := range nums {
		ans *= 10
		ans += num
	}
	return
}

func eval(b []rune) int {
	return 0
}
