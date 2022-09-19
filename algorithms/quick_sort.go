package algorithms

/*
Purpose:
Given an array of unsorted integers in an array, return an array of the same integers in ascending order by value.

Description:
Recursive / divide and conquer implementation.
Start by checking base case: does the array contain two or more comparable elements?
If it does not, return the array.

Otherwise, select a pivot element from the array and compare each element of the array to the pivot.
Group elements that are less than, equal to, and greater than the element.

Call QuickSort on the unsorted arrays (less, greater).
'Equal to' array is necessary to handle duplicate values but can be omitted from the implementation and replaced by the pivot if duplicates are not expected.

Return a new array which contains in order the contents of quicksorted 'less', 'equal to', and quicksorted 'greater'.

Worst Case Time Complexity:
O(n^2)

Pivot is leftmost element and array is sorted. As a result, 'less' is empty and 'equal to' and 'greater' become larger.
Unbalanced distribution between the arrays causes a longer total function call chain.

Average Case Time Complexity:
O(n log n) linearithmic

In the average case with a distributed and shuffled input, QuickSort will need to call itself log n times.
For each of those calls, it must make comparisons for the subset of n that is passed, resulting ~ n log n comparisons.

Space Complexity:
O(n) linear

This approach passes the input by value and creates new arrays for values less than, equal to, and greater than the pivot on each call to the function.
Since the new arrays are filled with the value of a QuickSort call and the full return array is only assembled once the base case is reached,
space complexity would appear to grow similarly to time complexity.

However, since the function creates a complete sorted copy of the input and does not sort in place, the space complexity must be at least equal to the input.
Therefore, space complexity scales at a linear rate with the input.
*/

func QuickSort(nums []int) []int {
	if len(nums) < 2 {

		return nums
	}

	less, equalto, greater := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, v := range nums {
		if v > nums[0] {
			greater = append(greater, v)
		}
		if v == nums[0] {
			equalto = append(equalto, v)
		}
		if v < nums[0] {
			less = append(less, v)
		}
	}

	less = QuickSort(less)
	greater = QuickSort(greater)

	return append(append(less, equalto...), greater...)
}
