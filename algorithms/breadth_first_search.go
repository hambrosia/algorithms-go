package algorithms

/*
Purpose:

Description:

Limitations:

Time Complexity:

Space Complexity:

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

	visited := make(map[int]bool)
	parents := make(map[int]int)

	found := false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		visited[node] = true

		if end == node {
			found = true
			break
		}

		for _, v := range graph[node] {
			if _, ok := visited[v]; !ok {
				parents[v] = node
				queue = append(queue, v)
			}
		}
	}

	if found {
		path := []int{end}
		for node, ok := parents[end]; ok; node, ok = parents[node] {
			path = append([]int{node}, path...)
		}
		return path
	}

	return []int{}
}
