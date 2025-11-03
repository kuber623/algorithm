package leetcode_656

import "math"

// https://leetcode.cn/problems/coin-path
// 难度：困难
// 题解：

func cheapestJump(coins []int, maxJump int) []int {
	// cost[i] 表示到达 i 时的最小成本
	cost := make([]int, len(coins))
	cost[0] = coins[0]
	// prev[i] 表示最小成本到达 i 的上一跳节点
	prev := make([]int, len(coins))
	prev[0] = -1

	// 使用单调队列获取滑动窗口区间的最小值
	queue := make([]int, 0, maxJump)
	for i := 0; i < len(cost)-1; i++ {
		for len(queue) > 0 && cost[tail(queue)] > cost[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		if front(queue) < i-maxJump+1 {
			queue = queue[1:]
		}

		// 如果滑动区间内所有节点都是不可达，表明不存在路径能到达终点
		// 故直接返回空数组
		if cost[front(queue)] == math.MaxInt {
			return []int{}
		}

		if coins[i+1] == -1 {
			cost[i+1] = math.MaxInt
		} else {
			cost[i+1] = cost[front(queue)] + cost[i+1]
			prev[i+1] = front(queue)
		}
	}

	path := make([]int, 0)
	for i := len(prev) - 1; i >= 0; {
		path = append(path, i)
		i = prev[i]
	}
	for i := 0; i <= len(path)/2; i++ {
		path[i], path[len(path)-1-i] = path[len(path)-1-i]+1, path[i]+1
	}
	return path
}

func front(queue []int) int {
	return queue[0]
}

func tail(queue []int) int {
	return queue[len(queue)-1]
}
