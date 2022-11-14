package problems

/*
Problem Definition:
Given a float representing a temperature in celsius, return an array of the same value in Kelvin and Fahrenheit.

Approach:
Do the math.

Worst Case Time Complexity:
Time complexity is constant, O(1).

Space Complexity:
Space complexity is always the amount of memory occupied by a slice of two float64 values. O(1)

With 'math-based' solutions, it's important to keep an eye out for overflow issues and set boundaries before performing calculations.
In this case, since the solutions uses a float data type, the number will lose precision rather than overflow as with an integer.
*/

func ConvertTemperature(celsius float64) []float64 {
	// return [kelvin, fahrenheit]
	return []float64{celsius + 273.15, (celsius * 1.8) + 32}
}
