package leetcode_3212

// https://leetcode.cn/problems/count-submatrices-with-equal-frequency-of-x-and-y/
// 难度：中等
// 题解：
// 用二维数组 sum 记录字符 X 和 Y 的个数

func numberOfSubmatrices(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])

	sum := make([][][2]int, m+1)
	for i := 0; i <= m; i++ {
		sum[i] = make([][2]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum[i+1][j+1][0] = sum[i+1][j][0] + sum[i][j+1][0] - sum[i][j][0]
			sum[i+1][j+1][1] = sum[i+1][j][1] + sum[i][j+1][1] - sum[i][j][1]
			switch grid[i][j] {
			case 'X':
				sum[i+1][j+1][0]++
			case 'Y':
				sum[i+1][j+1][1]++
			}
			if sum[i+1][j+1][0] == sum[i+1][j+1][1] && sum[i+1][j+1][0] > 0 {
				ans++
			}
		}
	}

	return
}
