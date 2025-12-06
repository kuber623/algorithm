package leetcode_1006

// https://leetcode.cn/problems/clumsy-factorial/
// 难度：中等

const (
	multiply = 0
	div      = 1
	add      = 2
	subtract = 3
)

func clumsy(n int) (ans int) {
	num, op := n, 0
	for i := n - 1; i > 0; i-- {
		switch op {
		case multiply:
			num = num * i
		case div:
			num = num / i
		case add:
			ans += num
			num = i
		case subtract:
			ans += num
			num = -i
		}
		op = (n - i) % 4
	}
	ans += num
	return
}
