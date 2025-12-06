package leetcode_456

// https://leetcode.cn/problems/132-pattern/
// 难度：中等
// 题解：根据枚举 1、3、2 分别有三种解法，这里只实现了枚举 2 解法

// 枚举 2 解法：遍历 2，并使用单调栈维护离 2 最近的 3，然后获取 3 之前数组的最小值是否大于 2
// 这个解法的好处就是无需提前了解数组全貌， 支持流式的数组输入
func find132patternEnum2(nums []int) bool {
	n := len(nums)

	// stack 为单调栈，用于寻找数组中每个元素上一个更大的元素（的索引号）
	// 将结果记录在 leftGreater 数组中
	stack, leftGreater := make([]int, 0, n), make([]int, n)
	for i := 0; i < n; i++ {
		// 将单调栈中小于等于 nums[i] 的元素出栈
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
			continue
		}

		if len(stack) > 0 {
			leftGreater[i] = stack[len(stack)-1]
		} else {
			leftGreater[i] = -1
		}

		stack = append(stack, i)
	}
	// minimum[i] 表示数组 [0, i] 区间的最小值
	minimum := make([]int, n)
	minimum[0] = nums[0]
	for i := 1; i < n; i++ {
		minimum[i] = min(minimum[i-1], nums[i])
	}

	for k := 2; k < n; k++ {
		j := leftGreater[k]
		if j < 1 {
			continue
		}
		if minimum[j-1] < nums[k] {
			return true
		}
	}

	return false
}
