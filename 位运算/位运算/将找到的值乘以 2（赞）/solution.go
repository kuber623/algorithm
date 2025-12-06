package leetcode_2154

// https://leetcode.cn/problems/keep-multiplying-found-values-by-two/
// 难度：简单

func findFinalValue(nums []int, original int) int {
	mask := 0
	for _, num := range nums {
		k := num / original
		if num%original == 0 && (k&(k-1)) == 0 {
			mask |= k
		}
	}

	return original * lowbit(^mask)
}

func lowbit(x int) int {
	return x & -x
}
