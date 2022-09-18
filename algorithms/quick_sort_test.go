package algorithms

import (
	"testing"

	. "github.com/hambrosia/agorithms-golang/testhelpers"
)

func TestQuickSort(t *testing.T) {
	// Values are out of order, no duplicates
	test_input := []int{10, 1, 30, 5, 50, 3}
	expected := []int{1, 3, 5, 10, 30, 50}

	result := QuickSort(test_input)
	AssertEquals(t, result, expected)

	// Values are all the same
	test_input = []int{11, 11, 11, 11, 11, 11}
	expected = []int{11, 11, 11, 11, 11, 11}

	result = QuickSort(test_input)
	AssertEquals(t, result, expected)

	// Values contain some duplicates
	test_input = []int{10, 1, 30, 5, 50, 3, 5, 10}
	expected = []int{1, 3, 5, 5, 10, 10, 30, 50}

	result = SelectionSort(test_input)
	AssertEquals(t, result, expected)

	// Values are alredy in order
	test_input = []int{1, 3, 5, 10, 30, 50}
	expected = []int{1, 3, 5, 10, 30, 50}

	result = QuickSort(test_input)
	AssertEquals(t, result, expected)
}
