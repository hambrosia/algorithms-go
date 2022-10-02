package algorithms

import (
	. "github.com/hambrosia/algorithms-golang/testhelpers"
	"testing"
)

func TestBreadthFirstSearchBool(t *testing.T) {

	// start, end, test graph (end value exists in graph)

	test_graph := make(map[int][]int)

	test_graph[0] = []int{1, 2, 3}
	test_graph[2] = []int{4, 5, 6}
	test_graph[3] = []int{7, 8}
	test_graph[6] = []int{9, 10}
	test_graph[10] = []int{13}

	test_start := 0
	test_end := 13

	res := BreadthFirstSearchBool(test_start, test_end, test_graph)

	AssertEquals(t, res, true)

	// start, end, test graph (end value not in graph)

	test_start = 0
	test_end = 12

	res = BreadthFirstSearchBool(test_start, test_end, test_graph)

	AssertEquals(t, res, false)
}

func TestBreadthFirstSearchPath(t *testing.T) {
	// start, end, test graph (end value exists in graph)

	test_graph := make(map[int][]int)

	test_graph[0] = []int{1, 2, 3}
	test_graph[2] = []int{4, 5, 6}
	test_graph[3] = []int{7, 8}
	test_graph[6] = []int{9, 10}
	test_graph[10] = []int{13}

	test_start := 0
	test_end := 13

	res := BreadthFirstSearchPath(test_start, test_end, test_graph)

	AssertEquals(t, res, []int{0, 2, 6, 10, 13})

	// start, end, test graph (end value not in graph)

	test_start = 0
	test_end = 12

	res = BreadthFirstSearchPath(test_start, test_end, test_graph)

	AssertEquals(t, res, []int{})

	// multiple paths to target value with different lengths
	tg2 := make(map[int][]int)

	tg2[0] = []int{1, 3, 5}
	tg2[1] = []int{6, 7, 8}
	tg2[5] = []int{9, 11, 13}
	tg2[7] = []int{9, 11}
	tg2[11] = []int{12}
	tg2[12] = []int{13}

	test_start = 0
	test_end = 13

	res = BreadthFirstSearchPath(test_start, test_end, tg2)

	AssertEquals(t, res, []int{0, 5, 13})

	// values of nodes in shortest path are not in order
	tg3 := make(map[int][]int)

	tg3[0] = []int{1, 3, 5}
	tg3[5] = []int{6, 7, 8}
	tg3[8] = []int{2, 4}
	tg3[2] = []int{9, 13}

	test_start = 0
	test_end = 13

	res = BreadthFirstSearchPath(test_start, test_end, tg3)

	AssertEquals(t, res, []int{0, 5, 8, 2, 13})

	// values of child nodes are not sorted

	tg4 := make(map[int][]int)

	tg4[0] = []int{1, 5, 2}
	tg4[5] = []int{9, 7, 10}
	tg4[9] = []int{11, 2, 10}
	tg4[10] = []int{11, 20, 13}
	tg4[9] = []int{13}
	tg4[2] = []int{13}

	test_start = 0
	test_end = 13

	res = BreadthFirstSearchPath(test_start, test_end, tg4)
	AssertEquals(t, res, []int{0, 2, 13})

}
