package algorithms

/*
Purpose:

Description:

Limitations:

Time Complexity:

Space Complexity:

*/

func BreadthFirstSearch(start int, end int, graph map[int][]int) bool {
	queue := make([]int, 0)
	queue = append(queue, graph[start]...)
	visited := make(map[int]bool)
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
				queue = append(queue, v)
			}
		}
	}
	return found
}
