package problems

import (
	"testing"

	. "github.com/hambrosia/algorithms-golang/testhelpers"
)

func TestConvertTemperature(t *testing.T) {

	// Example 1
	input := 36.5
	expected := []float64{309.65000, 97.70000}
	result := ConvertTemperature(input)
	AssertEquals(t, result, expected)

	// Example 2
	input = 122.11
	expected = []float64{395.26000, 251.79800}
	result = ConvertTemperature(input)
	AssertEquals(t, result, expected)
}
