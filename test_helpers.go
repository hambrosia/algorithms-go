package algorithms

import (
	"testing"
	"reflect"
)

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b){
		t.Errorf("%v is not equal to %v", a, b)
	}
}
