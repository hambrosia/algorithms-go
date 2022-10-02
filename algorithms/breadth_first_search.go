package algorithms

/*
Purpose:
Given a graph, determine if there is a path from one node to another.
If there is a path, return the shortest path between the node.

Description:
Visit the first node and mark it as visited.
For each neighbor node, mark it as visited. If it is not the target node, add its neighbor nodes to the list of nodes to check, and continue.

Limitations:
Doesn't work on weighted graphs. Doesn't work on undirected graphs.

Time Complexity:
O(V + E), in the worst case every vertex and edge must be traversed.


Space Complexity:
O(V), scales linearly with the number of vertices explored.
*/

func BreadthFirstSearchBool(start int, end int, graph map[int][]int) bool {
	queue := make([]int, 0)
	queue = append(queue, start)

	visited := make(map[int]bool)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		visited[node] = true
		if end == node {
			return true
		}

		for _, v := range graph[node] {
			if _, ok := visited[v]; !ok {
				queue = append(queue, v)
			}
		}
	}
	return false
}

func BreadthFirstSearchPath(start int, end int, graph map[int][]int) []int {
	queue := make([]int, 0)
	queue = append(queue, start)

	parents := make(map[int]int)
	parents[start] = -1

	found := false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if end == node {
			found = true
			break
		}

		for _, v := range QuickSort(graph[node]) {
			if _, ok := parents[v]; !ok {
				parents[v] = node
				queue = append(queue, v)
			}
		}
	}

	if found {
		path := []int{end}
		for node, ok := parents[end]; ok && node != -1; node, ok = parents[node] {
			path = append([]int{node}, path...)
		}
		return path
	}

	return []int{}
}
