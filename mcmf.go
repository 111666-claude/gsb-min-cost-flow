// Package mincostflow 求最小费用最大流。
// 缺陷：只增广一轮就返回，平行边互相覆盖，也没有势函数。
package mincostflow

import "errors"

// ErrInput 表示输入不合法。
var ErrInput = errors.New("mincostflow: 输入不合法")

// Edge 是一条有向边。
type Edge struct {
	From int
	To   int
	Cap  int64
	Cost int64
}

const infinity = int64(1) << 60

// Solve 返回 (最大流, 最小费用)。
func Solve(nodeCount int, edges []Edge, source, sink int) (int64, int64, error) {
	if nodeCount <= 0 || source < 0 || sink < 0 || source >= nodeCount || sink >= nodeCount {
		return 0, 0, ErrInput
	}
	capacity := map[[2]int]int64{}
	price := map[[2]int]int64{}
	for _, edge := range edges {
		if edge.From == edge.To || edge.Cap <= 0 {
			continue
		}
		if edge.From < 0 || edge.To < 0 || edge.From >= nodeCount || edge.To >= nodeCount {
			return 0, 0, ErrInput
		}
		capacity[[2]int{edge.From, edge.To}] = edge.Cap // 缺陷：覆盖而不是累加
		price[[2]int{edge.From, edge.To}] = edge.Cost
	}
	distance := make([]int64, nodeCount)
	parent := make([]int, nodeCount)
	for index := range distance {
		distance[index] = infinity
		parent[index] = -1
	}
	distance[source] = 0
	for round := 0; round < nodeCount; round++ {
		updated := false
		for edge, cap := range capacity {
			if cap <= 0 || distance[edge[0]] == infinity {
				continue
			}
			next := distance[edge[0]] + price[edge]
			if next < distance[edge[1]] {
				distance[edge[1]] = next
				parent[edge[1]] = edge[0]
				updated = true
			}
		}
		if !updated {
			break
		}
	}
	if distance[sink] == infinity {
		return 0, 0, nil
	}
	amount := infinity
	for node := sink; parent[node] >= 0; node = parent[node] {
		if capacity[[2]int{parent[node], node}] < amount {
			amount = capacity[[2]int{parent[node], node}]
		}
	}
	total := int64(0)
	for node := sink; parent[node] >= 0; node = parent[node] {
		total += amount * price[[2]int{parent[node], node}]
	}
	return amount, total, nil // 缺陷：只增广这一轮
}
