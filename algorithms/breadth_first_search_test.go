package algorithms

import (
	. "github.com/hambrosia/algorithms-golang/testhelpers"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {

	// start, end, test graph (end value exists in graph)

	test_graph := make(map[int][]int)

	test_graph[0] = []int{1, 2, 3}
	test_graph[2] = []int{4, 5, 6}
	test_graph[3] = []int{7, 8}
	test_graph[6] = []int{9, 10}
	test_graph[10] = []int{13}

	test_start := 0
	test_end := 13

	res := BreadthFirstSearch(test_start, test_end, test_graph)

	AssertEquals(t, res, true)

	// start, end, test graph (end value not in graph)

	test_start = 0
	test_end = 12

	res = BreadthFirstSearch(test_start, test_end, test_graph)

	AssertEquals(t, res, false)
}
