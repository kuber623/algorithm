package leetcode_2317

// https://leetcode.cn/problems/maximum-xor-after-operations/
// 难度：简单

func maximumXOR(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans |= num
	}
	return ans
}
