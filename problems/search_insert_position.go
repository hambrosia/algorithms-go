package problems

/*
Problem Definition:
Given an array of integers sorted by ascending value and containing no duplicates,
return the index of a target value.
If the target value is not found, return the index at which it should be inserted.


Approach:
This is a very light modification of binary search.
The main logic of the iterative binary search is not changed.
Instead, in the case that the element is not found, the return value is altered from -1 (not found),
to use information left over from the binary search to determine the index at which the value should be inserted.

The base case for the iterative binary search is when the low boundary of the examined sub-array becomes greater than the high boundary.

The last few steps before the base case is triggered help show why the low boundary will be the desired result
in the case that the value is not in the array.

When the value is not in the array, it is either less than the lowest value, greater than the highest value, or between
two values in the array.

In he final execution of the loop before the base case is triggered, the examined number input[mid] will be less than the target.
As a result, low will be set to mid + 1. On the next iteration of the loop, low will be greater
than high and therefore the loop will not run and low will be returned.


Worst Case Time Complexity:
O(log n)
It's binary search; the number of compares needed is cut by half on each iteration.
As a result it scales in base2 log with the size of the input.

Space Complexity:
O(1)
This can be accomplished without copying the input array. However, this implementation
passes by value and thus the function duplicates the input in memory.
*/

func SearchInsertPosition(nums []int, target int) int {
	high_bound := len(nums) - 1
	low_bound := 0
	for high_bound >= low_bound {
		mid := (high_bound + low_bound) / 2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low_bound = mid + 1
		} else if nums[mid] > target {
			high_bound = mid - 1
		}
	}

	return low_bound
}
