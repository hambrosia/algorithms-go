package testhelpers

import (
	"reflect"
	"testing"
)

func AssertEquals(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%v is not equal to %v", a, b)
	}
}
