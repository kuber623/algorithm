package leetcode_239

func maxSlidingWindow(nums []int, k int) []int {
	// 处理边界情况
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	queue := make([]int, 0, k+1) // 单调递减队列
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		// 如果队尾元素小于当前元素值，则弹出队尾
		for len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
			continue
		}

		queue = append(queue, i)

		// 如果队首元素所在位置超出窗口位置，则弹出队首
		if queue[0] < i-k+1 {
			queue = queue[1:]
		}

		ans[i] = nums[queue[0]]
	}

	return ans[k-1:]
}
