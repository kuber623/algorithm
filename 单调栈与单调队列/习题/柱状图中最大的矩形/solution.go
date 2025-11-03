package leetcode_84

import "math"

// https://leetcode.cn/problems/largest-rectangle-in-histogram/
// 难度：困难
// 题解：单调栈维护右边界

func largestRectangleArea(heights []int) int {
	n := len(heights)

	lesser := nextLesserElement(heights)

	area := 0
	for i := 0; i < n; i++ {
		for j := i; j != math.MaxInt; {
			height := heights[j]
			width := min(lesser[j]-i, len(heights)-i)
			area = max(area, height*width)
			j = lesser[j]
		}
	}

	return area
}

func nextLesserElement(nums []int) []int {
	lesser := make([]int, len(nums))

	stack := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
			continue
		}

		if len(stack) == 0 {
			lesser[i] = math.MaxInt
		} else {
			lesser[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	return lesser
}
