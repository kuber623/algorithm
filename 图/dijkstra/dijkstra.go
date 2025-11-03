package dijkstra

import "math"

// Dijkstra 算法是一种经典的「单源最短路径」算法，它的核心作用是在带非负权重的图中，从一个指定的源顶点出发，计算该顶点到图中所有其他顶点的最短路径长度
// 其核心思想是：
// - 基于「贪心策略」每次从「未确定最短路径的顶点」中选择当前距离源顶点最近的顶点，确定其最短路径
// - 然后以该顶点为「中间点」，更新其他未确定顶点的距离
//
// - 时间复杂度：O(n²)
// - 空间复杂度：O(n)
func Dijkstra(graph [][]int, start int) []int {
	n := len(graph)
	distance := make([]int, n) // distance 数组记录目标节点到当前节点的已知最短距离
	finial := make([]bool, n)  // finial 数组记录目标节点到当前节点的最短距离是否已被找到
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt
	}

	cur := start
	for {
		for i := range graph[cur] {
			if finial[i] || graph[cur][i] == math.MaxInt {
				continue
			}
			dist := distance[cur] + graph[cur][i]
			if dist < distance[i] {
				distance[i] = dist
			}
		}

		minDist := math.MaxInt
		for i := 0; i < n; i++ {
			if finial[i] {
				continue
			}
			if distance[i] <= minDist {
				minDist, cur = distance[i], i
			}
		}
		finial[cur] = true

		if minDist == math.MaxInt {
			break
		}
	}

	return distance
}
