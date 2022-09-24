package problems

/*
Problem Definition:
Given an array containing y-coordinate values (height) as integers, and treating index values as x-coordinate values,
return the integer value of the largest rectangle combining two 'lines' in the array.

In the array [1,8,6,2,5,4,8,3,7], indexes 1 and 8 would form the largest rectangle with an area of 49.
If the lines are of differing heights, the rectangle must be constructed with the lesser of the heights of the two ines.
The width is the difference between the indexes of the values.

Approach:
The 'easy' approach to this problem is to nest loops and iterate over the rest of the array from each index examined.
This results in O(n^2) time complexity!

Another simple approach is to maintain two boundaries and compare opposite ends of the sub-array while decreasing its size
on each iteration. This will result in O(n) time complexity and is the approach used below.

There may be a solution involving sorting, since the maximum rectangle is usually the result of a value that has a low index
and large height with a value that has a high index and high height. I couldn't get this working though.

Worst Case Time Complexity:
The two-boundary approach has O(n) time complexity since the array is traversed once.

Space Complexity:
Either O(1) or O(n) depending on if the input is passed by reference or by value.
*/

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func ContainerWithMostWater(height []int) int {
	left := 0
	right := len(height) - 1
	max_rect := -1
	for left < right {
		left_h := height[left]
		right_h := height[right]

		h := min(left_h, right_h)
		w := right - left

		max_rect = max(max_rect, h*w)

		if right_h > left_h {
			left++
		} else {
			right--
		}
	}
	return max_rect
}
