package problems

import (
	"fmt"
	"math"
)

func AddBinary(a string, b string) string {

	var carry uint8 = 0
	result := ""
	for i := 0; i < int(math.Max(float64(len(a)), float64(len(b)))) || carry > 0; i++ {
		var a_val uint8 = 0
		if i < len(a) {
			a_val = uint8(a[len(a)-i-1] - 48) // ASCII codes for numbers start at 48, subtract 48 to start at 0
		}

		var b_val uint8 = 0
		if i < len(b) {
			b_val = uint8(b[len(b)-i-1] - 48)
		}

		sum := a_val + b_val + carry
		fmt.Printf("old result: %s\n", result)
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
