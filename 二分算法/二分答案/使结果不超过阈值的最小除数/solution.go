package leetcode_1283

import (
	"slices"
)

// https://leetcode.cn/problems/find-the-smallest-divisor-given-a-threshold/
// 难度：中等
// 题解：
// 当被除数为 1 时，和值为最大值。随着被除数的递增，和值会逐渐下降。直到被除数为数组的最大值，此时和值达到最小值，即数组长度
// 因此我们可以设置被除数的区间为 [1, MAX{nums}]，然后通过二分算法找到第一个使结果不超过阈值的最小除数
// 另外需要注意的是，这里的除法需要对结果向上取整，因此执行除法预算时可以对被除数增加 divisor + 1 后，对结果向下取整

func smallestDivisor(nums []int, threshold int) int {
	left, right := 1, slices.Max(nums)
	for left <= right {
		mid := left + (right-left)>>1

		sum := 0
		for _, n := range nums {
			sum += (n + mid - 1) / mid
		}

		if sum > threshold {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
