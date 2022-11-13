package problems

import (
	"math"
)

/*
Problem Definition:
Given two binary numbers represented as a string, return the sum of the numbers as a binary number string.

Approach:
Iterate from right to left and sum each digit of the binary number, using a carry to pass overflow to the next
digit to the left. Continue until we've reached the maximum length of either of the two original numbers and the
carry is empty.

Worst Case Time Complexity:
Linear O(n) time complexity is guaranteed since this approach always iterates through the length of the longest
number. As an optimization, the algorithm could exit early if one number is shorter than the other and the carry
is empty. This would still result in O(n) worst case time complexity, but could result in lower times for some
numbers.

Space Complexity:
Roughly linear / O(n)
The result is at least as large as the greatest of the two inputs. Additional space is also used for the carry
and for intermediate sums.
*/

func AddBinary(a string, b string) string {

	var carry uint8 = 0
	result := ""
	for i := 0; i < int(math.Max(float64(len(a)), float64(len(b)))) || carry > 0; i++ {

		var a_val uint8 = 0
		if i < len(a) { // safe get from each array and default to 0 if element out of bounds
			a_val = uint8(a[len(a)-i-1] - 48) // ASCII codes for numbers start at 48, subtract 48 to start at 0
		}

		var b_val uint8 = 0
		if i < len(b) {
			b_val = uint8(b[len(b)-i-1] - 48)
		}

		sum := a_val + b_val + carry
		switch sum {
		case 0:
			carry = 0
			result = "0" + result
		case 1:
			carry = 0
			result = "1" + result
		case 2:
			carry = 1
			result = "0" + result
		case 3:
			carry = 1
			result = "1" + result
		}
	}
	return result
}
