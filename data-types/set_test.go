package datatypes

import (
	"testing"

	. "github.com/hambrosia/algorithms-golang/testhelpers"
)

func TestSetBasicReadWriteMethods(t *testing.T) {

	s1 := Constructor()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("baz")
	AssertEquals(t, 3, s1.Len())
	AssertEquals(t, true, s1.Contains("foo"))

	s1.Remove("foo")
	AssertEquals(t, 2, s1.Len())

	s1.Clear()
	AssertEquals(t, 0, s1.Len())
	AssertEquals(t, true, s1.IsEmpty())
}

func TestDiff(t *testing.T) {
	s1 := Constructor()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("baz")

	s2 := Constructor()
	s2.Add("foo")
	s2.Add("bar")

	diff := s1.Difference(&s2)
	AssertEquals(t, 1, diff.Len())
	AssertEquals(t, true, diff.Contains("baz"))
}

func TestSymmetricDifference(t *testing.T) {
	s1 := Constructor()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("baz")

	s2 := Constructor()
	s2.Add("foo")
	s2.Add("bar")
	s2.Add("banana")

	diff := s1.SymmetricDifference(&s2)
	AssertEquals(t, 2, diff.Len())
	AssertEquals(t, true, diff.Contains("baz"))
	AssertEquals(t, true, diff.Contains("banana"))

}

func TestIntersection(t *testing.T) {
	s1 := Constructor()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("baz")

	s2 := Constructor()
	s2.Add("foo")
	s2.Add("bar")
	s2.Add("banana")

	diff := s1.Intersection(&s2)
	AssertEquals(t, 2, diff.Len())
	AssertEquals(t, true, diff.Contains("foo"))
	AssertEquals(t, true, diff.Contains("bar"))
}

func TestUnion(t *testing.T) {
	s1 := Constructor()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("baz")

	s2 := Constructor()
	s2.Add("foo")
	s2.Add("bar")
	s2.Add("banana")

	diff := s1.Union(&s2)
	AssertEquals(t, 4, diff.Len())
	AssertEquals(t, true, diff.Contains("foo"))
	AssertEquals(t, true, diff.Contains("bar"))
	AssertEquals(t, true, diff.Contains("baz"))
	AssertEquals(t, true, diff.Contains("banana"))

}

func TestIsSuperSet(t *testing.T) {
	s1 := Constructor()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("baz")

	s2 := Constructor()
	s2.Add("foo")
	s2.Add("bar")
	s2.Add("banana")

	AssertEquals(t, false, s2.IsSuperset(&s1))

	s1.Remove("baz")
	AssertEquals(t, true, s2.IsSuperset(&s1))

}

func TestIsSubset(t *testing.T) {
	s1 := Constructor()
	s1.Add("foo")
	s1.Add("bar")
	s1.Add("baz")

	s2 := Constructor()
	s2.Add("foo")
	s2.Add("bar")
	s2.Add("banana")

	AssertEquals(t, false, s1.IsSubset(&s2))

	s1.Remove("baz")
	AssertEquals(t, true, s1.IsSubset(&s2))
}

// func TestDifferenceUpdate(t *testing.T){
// 	s1 := Constructor()
// 	s1.Add("foo")
// 	s1.Add("bar")
// 	s1.Add("baz")

// 	s2 := Constructor()
// 	s2.Add("foo")
// 	s2.Add("bar")
// 	s2.Add("banana")

// 	s1.DifferenceUpdate(&s2)
// 	AssertEquals(t, 1, s1.Len())
// 	AssertEquals(t, true, s1.Contains("baz"))
// }
