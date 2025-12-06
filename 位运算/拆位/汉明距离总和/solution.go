package leetcode_477

// https://leetcode.cn/problems/total-hamming-distance/
// 难度：中等

func totalHammingDistance(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < 31; i++ {
		cnt := 0
		for _, num := range nums {
			if (num>>i)&1 == 0 {
				cnt++
			}
		}
		ans += cnt * (n - cnt)
	}
	return
}
