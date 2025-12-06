package leetcode_798

// https://leetcode.cn/problems/smallest-rotation-with-highest-score/
// 难度：困难

func bestRotation(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		nums[i] = nums[i] - i
	}

	score := make([]int, n+1)
	for i := 0; i < n; i++ {
		if i < n-nums[i] {
			score[i+1]++
			score[n-nums[i]+1]--
		}

		b := min(0, i+1, -nums[i]+1)
		score[0]++
		score[b]--
	}

	sum, best, k := 0, 0, 0
	for i := 0; i < n; i++ {
		sum += score[i]
		if sum > best {
			best = sum
			k = i
		}
	}

	return k
}
