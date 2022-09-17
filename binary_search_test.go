package algorithms

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	test_input := make([]int, 0)
	test_target := 0
	result := binarySearch(&test_input, &test_target)

	// List is empty, target is 0
	expected := -1
	assertEquals(t, result, expected)

	// Value is in array
	test_input = []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	test_target = 7
	expected = 4

	result = binarySearch(&test_input, &test_target)
	assertEquals(t, result, expected)

	// Value is not in array
	test_input = []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	test_target = 27
	expected = -1

	result = binarySearch(&test_input, &test_target)
	assertEquals(t, result, expected)
}
