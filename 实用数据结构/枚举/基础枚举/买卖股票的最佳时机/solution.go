package leetcode_121

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 难度：简单
// 题解：
// 第 i 天能够过去到的最大收益是 prices[i] - cost 其中 cost 为前 i-1 天的最低买入价格
// 因此可以通过一次遍历即可获取到最大利润

func maxProfit(prices []int) int {
	cost, profit := math.MaxInt, 0
	for _, price := range prices {
		profit = max(profit, price-cost)
		cost = min(cost, price)
	}
	return profit
}
