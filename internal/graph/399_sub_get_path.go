package graph

// 除法求值
// 给定除法数组和值数组，求指定除法的结果
// 图的连接算法 floyd算法求每个节点之间的路径连乘值
func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	// 建图
	graph := make([][]float64, len(id))
	for i := range graph {
		graph[i] = make([]float64, len(id))
	}
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		graph[v][w] = values[i]
		graph[w][v] = 1 / values[i]
	}

	// 执行 Floyd 算法
	for k := range graph {
		for i := range graph {
			for j := range graph {
				if graph[i][k] > 0 && graph[k][j] > 0 {
					graph[i][j] = graph[i][k] * graph[k][j]
				}
			}
		}
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if !hasS || !hasE || graph[start][end] == 0 {
			ans[i] = -1
		} else {
			ans[i] = graph[start][end]
		}
	}
	return ans
}
