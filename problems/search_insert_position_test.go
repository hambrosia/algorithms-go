package problems

import (
	"testing"

	. "github.com/hambrosia/algorithms-golang/testhelpers"
)

func TestSearchInsertPosition(t *testing.T) {
	test_input := []int{1, 3, 5, 6}
	test_target := 2
	expected := 1
	res := SearchInsertPosition(test_input, test_target)
	AssertEquals(t, res, expected)

	test_input = []int{1, 3, 5, 6}
	test_target = 5
	expected = 2
	res = SearchInsertPosition(test_input, test_target)
	AssertEquals(t, res, expected)

	test_input = []int{1, 3, 5, 6}
	test_target = 7
	expected = 4
	res = SearchInsertPosition(test_input, test_target)
	AssertEquals(t, res, expected)

	test_input = []int{1, 3, 5, 6}
	test_target = 0
	expected = 0
	res = SearchInsertPosition(test_input, test_target)
	AssertEquals(t, res, expected)

	test_input = []int{1, 3}
	test_target = 2
	expected = 1
	res = SearchInsertPosition(test_input, test_target)
	AssertEquals(t, res, expected)

	test_input = []int{3, 5, 7, 8, 9}
	test_target = 1
	expected = 0
	res = SearchInsertPosition(test_input, test_target)
	AssertEquals(t, res, expected)

	test_input = []int{3, 5, 7, 8, 9}
	test_target = 12
	expected = 5
	res = SearchInsertPosition(test_input, test_target)
	AssertEquals(t, res, expected)

}
