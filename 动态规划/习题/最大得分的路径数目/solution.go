package leetcode_1301

// https://leetcode.cn/problems/number-of-paths-with-max-score/
// 难度：困难

func pathsWithMaxScore(board []string) []int {
	m, n := len(board), len(board[0])

	// score[i][j] 表示从 (i, j) 到 (m-1, n-1) 的最大得分
	score := make([][]int, m+1)
	for i := range score {
		score[i] = make([]int, n+1)
	}
	// path[i][j] 代表从 (i, j) 到 (m-1, n-1) 获得最大得分的路径数
	path := make([][]int, m+1)
	for i := range path {
		path[i] = make([]int, n+1)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			switch board[i][j] {
			case 'S':
				score[i][j], path[i][j] = 0, 1
			case 'E':
				if noPath(path, i, j) {
					continue
				}
				// 计算得分
				score[i][j] = max(score[i+1][j], score[i][j+1], score[i+1][j+1])
				// 计算路径数
				path[i][j] = getPath(path, score, i, j)
			case 'X':
				continue
			default:
				if noPath(path, i, j) {
					continue
				}
				// 计算得分
				score[i][j] = max(score[i+1][j], score[i][j+1], score[i+1][j+1]) + int(board[i][j]-'0')
				// 计算路径数
				path[i][j] = getPath(path, score, i, j)
			}
		}
	}

	return []int{score[0][0], path[0][0]}
}

func noPath(path [][]int, i, j int) bool {
	return path[i+1][j] == 0 && path[i][j+1] == 0 && path[i+1][j+1] == 0
}

func getPath(path [][]int, score [][]int, i, j int) int {
	count := 0
	switch {
	case score[i+1][j] > score[i][j+1]:
		count = path[i+1][j]
	case score[i+1][j] < score[i][j+1]:
		count = path[i][j+1]
	default:
		if path[i+1][j] == 0 && path[i][j+1] == 0 {
			count = path[i+1][j+1]
		} else {
			count = path[i+1][j] + path[i][j+1]
		}
	}

	return count % 1000000007
}
