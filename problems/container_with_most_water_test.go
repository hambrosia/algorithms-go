package problems

import (
	. "github.com/hambrosia/agorithms-golang/testhelpers"
	"testing"
)

func TestContainerWithMostWater(t *testing.T) {

	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expected := 49
	res := ContainerWithMostWater(input)
	AssertEquals(t, res, expected)

	input = []int{1, 1}
	expected = 1
	res = ContainerWithMostWater(input)
	AssertEquals(t, res, expected)
}
