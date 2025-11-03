package leetcode_96

// https://leetcode.cn/problems/unique-binary-search-trees
// 难度：中等
// 题解：
// 二叉搜索树的种数可以通过枚举树根值 1 ~ n 累加得到
// 对于每种情况，其二叉树种数为「左子树种数 * 右子树种数」
// 由于二叉搜索树的种数仅于树的节点数量有关，因此可以通过记录每种节点数量的二叉搜索树种数

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}
	}

	return dp[n]
}
