package algorithms

import (
	"testing"
)

func assertEquals(t *testing.T, a int, b int) {
	if a != b {
		t.Errorf("Integer %v is not equal to %v", a, b)
	}
}
