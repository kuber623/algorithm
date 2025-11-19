package leetcode_2848

// https://leetcode.cn/problems/points-that-intersect-with-cars/
// 难度：简单
// 题解：差分

func numberOfPoints(nums [][]int) int {
	diff := make([]int, 101)

	for _, num := range nums {
		begin, end := num[0], num[1]
		diff[begin]++
		diff[end+1]--
	}

	ans, sum := 0, 0
	for i := 0; i < len(diff); i++ {
		sum += diff[i]
		if sum > 0 {
			ans++
		}
	}

	return ans
}
