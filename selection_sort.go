package algorithms

/*
Purpose:
Given an array of unsorted integers in an array, return an array of the same integers in ascending order by value.

Description:
Prepare an empty array for the sorted result.
Remove the smallest value from the original array and append it to the new array.
Repeat until original array is empty.
Return the new array.


Worst Case Time Complexity:
O(n^2)

The list is sorted in descending order. As a result, the implementation must examine every value of the array on each pass in order to select the smallest value.

Space Complexity:
O(n)

This implementation does not sort in place and returns a new array.

*/

func selectionSort(nums []int) []int {
	ret := make([]int, 0)

	for len(nums) > 0 {
		index_smallest := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] < nums[index_smallest] {
				index_smallest = i
			}
		}
		ret = append(ret, nums[index_smallest])
		nums[index_smallest] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
	}

	return ret
}
