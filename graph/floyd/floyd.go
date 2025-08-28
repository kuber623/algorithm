package floyd

import "math"

// Floyd 弗洛伊德算法是一种用于求解图中「所有顶点之间最短路径」的动态规划算法，它适用于有向图或无向图，支持带权边
// 其核心思想如下：
// 1. 初始时直接记录顶点之间的直达距离（使用 math.MaxInt 表示两点不直接相连）
// 2. 引入中间顶点 K，尝试通过中间节点 K 缩短 I 和 J 两点间距离：如果 I -> K -> J 比 I -> J 的路径要短，则更新 I -> J 的最短距离
// 3. 在完成所有节点的迭代后，得到所有顶点间的最短路径
//
// - 时间复杂度：O(n³)
// - 空间复杂度：O(n²)
func Floyd(graph [][]int) [][]int {
	n := len(graph)
	distances := make([][]int, n)
	for i := range distances {
		copy(distances[i], graph[i])
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distances[i][k] == math.MaxInt || distances[k][j] == math.MaxInt {
					continue
				}
				d := distances[i][k] + distances[k][j]
				if d < distances[i][j] {
					distances[i][j] = d
				}
			}
		}
	}

	return distances
}
