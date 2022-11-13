package problems

import (
	"testing"

	. "github.com/hambrosia/algorithms-golang/testhelpers"
)

func TestAddBinary(t *testing.T) {

	// example 1
	a := "11"
	b := "1"
	expected := "100"
	result := AddBinary(a, b)
	AssertEquals(t, result, expected)

	// example 2
	a = "1010"
	b = "1011"
	expected = "10101"
	result = AddBinary(a, b)
	AssertEquals(t, result, expected)
}
