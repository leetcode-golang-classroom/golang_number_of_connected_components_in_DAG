package sol

type nodes []int

func countNumberOfConnectedComponents(n int, edges [][]int) int {
	adjacencyMap := make(map[int]nodes, n)
	for _, edge := range edges {
		adjacencyMap[edge[0]] = append(adjacencyMap[edge[0]], edge[1])
		adjacencyMap[edge[1]] = append(adjacencyMap[edge[1]], edge[0])
	}
	visit := make(map[int]struct{})
	result := 0
	var dfs func(node int)
	dfs = func(node int) {
		if _, ok := visit[node]; ok {
			return
		}
		visit[node] = struct{}{}
		adjacencyNodes := adjacencyMap[node]
		for _, adjacentNode := range adjacencyNodes {
			if _, visited := visit[adjacentNode]; !visited {
				dfs(adjacentNode)
			}
		}
	}
	for node := 0; node < n; node++ {
		_, visited := visit[node]
		if !visited {
			dfs(node)
			result++
		}
	}
	return result
}
