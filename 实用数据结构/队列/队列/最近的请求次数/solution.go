package leetcode_933

// https://leetcode.cn/problems/number-of-recent-calls/
// 难度：简单

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (counter *RecentCounter) Ping(t int) int {
	counter.queue = append(counter.queue, t)

	i := 0
	for _, ot := range counter.queue {
		if ot >= t-3000 {
			break
		}
		i++
	}
	counter.queue = counter.queue[i:]

	return len(counter.queue)
}
