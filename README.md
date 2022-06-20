# golang_number_of_connected_components_in_DAG

Given n nodes labeled from 0 to n - 1 and a list of undirected edges (each edge is a pair of nodes), write a function to find the number of connected components in an undirected graph.

## Examples

**Example 1:**

```
     0          3
     |          |
     1 --- 2    4
Given n = 5 and edges = [[0, 1], [1, 2], [3, 4]], return 2.
```

**Example 2:**

```
     0           4
     |           |
     1 --- 2 --- 3
Given n = 5 and edges = [[0, 1], [1, 2], [2, 3], [3, 4]], return 1.
```

**Note:**
You can assume that no duplicate edges will appear in edges. Since all edges are undirected, [0, 1] is the same as [1, 0] and thus will not appear together in edges.

## 解析

題目給定一個整數 n 代表有 0 到 n-1 vertex

還有一個 edges 矩陣用來表示有相連的 edge

要求寫一個演算法來計算有相連的區塊個數

作法1:

透過 edges 可以推算出 adjacency list 

然後針對 adjacency list 做 DFS 走訪所有相連的區塊

![](https://i.imgur.com/ChnlqUm.png)

這樣的時間複雜度 O(e+v) 其中 e 代表 edge 個數, v 代表 vertex 個數

作法 2:

 UnionFind 作法

初始化每個 vertex 為自己的 parent

初始化每個 vertex rank 為 1

逐步把每個 edge的兩個 vertex 做 union

如果 union 成功則返回 1 , 如果沒有 union 則返回 0

所以把 n - 每次 union 結果就是答案

![](https://i.imgur.com/E9s7zex.png)

這樣只要走完所有 edge

所時間複雜度為 O(e)

## 程式碼
```go
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

```
## 困難點

1. 理解如何透過 adjacency list 做累計

## Solve Point

- [x]  透過 edge 資訊找出 adjacency list
- [x]  透過 DFS 來走訪相連的區塊