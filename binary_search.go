package algorithms

/*
Purpose:
Given a list of integers sorted in ascedning value and a target integer, return the index of the target integer, or -1 if the integer is not in the list.

Description:
Set lower and upper boundaries to the full size of the input array (first and last indexes).
Run a loop so long as the lower boundary is less than or equal to the higher boundary.
The condition for the loop is the base case of Binary Search.
Binary Search shrinks the examined slice of the array by half until it can shrink no more or the value is found.
Within the loop, first set the middle value to the value of the lower boundary plus the upper boundary and divide in half.
Compare the middle value to the target value.
If the target is found, return it.
If the middle value is less than the target, discard all values that are less than the middle value (set new lower bound to mid plus one).
If the middle value is greater than the target, discard all values that are greater than the middle value (set new higher bound to minus one).
The new boundaries created are padded by one in the appropriate direction since the middle value has already been comapred to the target.
The loops continues until the base case is satisfied (target found and returned, or boundaries meet and loop will not run).
If the loop finishes without returning the value, the value was not found and Binary Search indicates this by returning -1.

Limitations:
This implementation does not handle duplicate adjacent values. A modification that could handle that is to check adjacent values of mid as part of the base case test.

Worst Case Time Complexity:
O(log n)

The target is in the first or last index of the input. Log n comparisons must be made.


Space Complexity:
O(1)

This implementation maintains integers for the lower and upper boundaries, and the middle value.
The input array and target are passed by reference to avoid consuming additional memory.
*/

func binarySearch(nums *[]int, target *int) int {
	high_bound := len(*nums) - 1
	low_bound := 0

	for low_bound <= high_bound {
		mid := (low_bound + high_bound) / 2
		if (*nums)[mid] == *target {
			return mid
		}
		if (*nums)[mid] < *target {
			low_bound = mid + 1
		}
		if (*nums)[mid] > *target {
			high_bound = mid - 1
		}
	}
	return -1
}
