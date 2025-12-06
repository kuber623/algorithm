package leetcode_1526

// https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
// 难度：困难
// 题解：

func minNumberOperations(target []int) int {
	n := len(target)

	// 通过单调栈获取数组中每个元素右侧首个更小元素
	smaller, stack := make([]int, n), make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		num := target[i]

		for len(stack) > 0 && num <= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			smaller[i] = stack[len(stack)-1]
		} else {
			smaller[i] = -1
		}

		stack = append(stack, num)
	}

	for i := 0; i < n; i++ {

	}

	return 0
}
