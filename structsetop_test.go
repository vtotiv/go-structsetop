package structsetop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// A few structs to test on
type A struct {
	va int
}
type B struct {
	vb int
}
type C struct {
	vc int
}

func TestUnion(t *testing.T) {

}

func TestIntersection(t *testing.T) {
	arrOfA := []A{{1}, {2}, {3}}
	arrOfB := []B{{2}, {3}, {4}}
	arrOfC := []C{{4}, {5}}
	arrOfB2 := []B{}

	expectIntersectionAB := []A{{2}, {3}}
	expectIntersectionAC := []A{}
	expectIntersectionBC := []B{{4}}
	expectIntersectionAB2 := []A{}

	equalAB := func(a *A, b *B) bool {
		return a.va == b.vb
	}
	equalAC := func(a *A, c *C) bool {
		return a.va == c.vc
	}
	equalBC := func(b *B, c *C) bool {
		return b.vb == c.vc
	}

	assert.Equal(t, expectIntersectionAB, Intersection(arrOfA, arrOfB, equalAB))
	assert.Equal(t, expectIntersectionAC, Intersection(arrOfA, arrOfC, equalAC))
	assert.Equal(t, expectIntersectionBC, Intersection(arrOfB, arrOfC, equalBC))
	assert.Equal(t, expectIntersectionAB2, Intersection(arrOfA, arrOfB2, equalAB))
}

func TestDifference(t *testing.T) {

}

func TestSymmetricDifference(t *testing.T) {

}

func TestSubset(t *testing.T) {

}

func TestSuperset(t *testing.T) {

}
