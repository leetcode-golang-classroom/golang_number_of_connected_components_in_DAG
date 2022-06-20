package sol

func countNumberOfConnectedComponentsV1(n int, edges [][]int) int {
	parent := make([]int, n)
	rank := make([]int, n)
	for node := 0; node < n; node++ {
		parent[node] = node
		rank[node] = 1
	}
	var find func(node int) int
	find = func(node int) int {
		p := parent[node]
		for p != parent[p] {
			parent[p] = parent[parent[p]]
			p = parent[p]
		}
		return p
	}
	var union func(node1, node2 int) int
	union = func(node1, node2 int) int {
		p1 := find(node1)
		p2 := find(node2)
		if p1 == p2 { // not union
			return 0
		}
		if rank[p1] > rank[p2] {
			parent[p2] = p1
			rank[p1] += rank[p2]
		} else {
			parent[p1] = p2
			rank[p2] += rank[p1]
		}
		return 1
	}
	result := n
	for _, edge := range edges {
		result -= union(edge[0], edge[1])
	}
	return result
}
