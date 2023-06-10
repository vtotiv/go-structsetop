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
	arrOfA := []A{{1}, {2}, {3}}
	arrOfB := []B{{2}, {3}, {4}}
	arrOfC := []C{{4}, {5}}
	arrOfA2 := []A{}
	arrOfB2 := []B{}

	expectDifferenceAB := []A{{1}}
	expectDifferenceAC := []A{{1}, {2}, {3}}
	expectDifferenceBC := []B{{2}, {3}}
	expectDifferenceA2B := []A{}
	expectDifferenceAB2 := []A{{1}, {2}, {3}}

	equalAB := func(a *A, b *B) bool {
		return a.va == b.vb
	}
	equalAC := func(a *A, c *C) bool {
		return a.va == c.vc
	}
	equalBC := func(b *B, c *C) bool {
		return b.vb == c.vc
	}

	assert.Equal(t, expectDifferenceAB, Difference(arrOfA, arrOfB, equalAB))
	assert.Equal(t, expectDifferenceAC, Difference(arrOfA, arrOfC, equalAC))
	assert.Equal(t, expectDifferenceBC, Difference(arrOfB, arrOfC, equalBC))
	assert.Equal(t, expectDifferenceA2B, Difference(arrOfA2, arrOfB, equalAB))
	assert.Equal(t, expectDifferenceAB2, Difference(arrOfA, arrOfB2, equalAB))
}

func TestSymmetricDifference(t *testing.T) {
	arrOfA1 := []A{{1}, {2}, {3}}
	arrOfA2 := []A{}
	arrOfB1 := []B{{2}, {3}, {4}}
	arrOfB2 := []B{}
	arrOfB3 := []B{{3}, {4}, {5}}
	arrOfB4 := []B{{4}, {5}, {6}}

	expectSymmetricDifferenceA1B1A := []A{{1}}
	expectSymmetricDifferenceA1B1B := []B{{4}}
	expectSymmetricDifferenceA1B2A := []A{{1}, {2}, {3}}
	expectSymmetricDifferenceA1B2B := []B{}
	expectSymmetricDifferenceA2B1A := []A{}
	expectSymmetricDifferenceA2B1B := []B{{2}, {3}, {4}}
	expectSymmetricDifferenceA2B2A := []A{}
	expectSymmetricDifferenceA2B2B := []B{}
	expectSymmetricDifferenceA1B3A := []A{{1}, {2}}
	expectSymmetricDifferenceA1B3B := []B{{4}, {5}}
	expectSymmetricDifferenceA1B4A := []A{{1}, {2}, {3}}
	expectSymmetricDifferenceA1B4B := []B{{4}, {5}, {6}}

	equalAB := func(a *A, b *B) bool {
		return a.va == b.vb
	}

	a, b := SymmetricDifference(arrOfA1, arrOfB1, equalAB)
	assert.Equal(t, expectSymmetricDifferenceA1B1A, a)
	assert.Equal(t, expectSymmetricDifferenceA1B1B, b)

	a, b = SymmetricDifference(arrOfA1, arrOfB2, equalAB)
	assert.Equal(t, expectSymmetricDifferenceA1B2A, a)
	assert.Equal(t, expectSymmetricDifferenceA1B2B, b)

	a, b = SymmetricDifference(arrOfA2, arrOfB1, equalAB)
	assert.Equal(t, expectSymmetricDifferenceA2B1A, a)
	assert.Equal(t, expectSymmetricDifferenceA2B1B, b)

	a, b = SymmetricDifference(arrOfA2, arrOfB2, equalAB)
	assert.Equal(t, expectSymmetricDifferenceA2B2A, a)
	assert.Equal(t, expectSymmetricDifferenceA2B2B, b)

	a, b = SymmetricDifference(arrOfA1, arrOfB3, equalAB)
	assert.Equal(t, expectSymmetricDifferenceA1B3A, a)
	assert.Equal(t, expectSymmetricDifferenceA1B3B, b)

	a, b = SymmetricDifference(arrOfA1, arrOfB4, equalAB)
	assert.Equal(t, expectSymmetricDifferenceA1B4A, a)
	assert.Equal(t, expectSymmetricDifferenceA1B4B, b)
}

func TestSubset(t *testing.T) {
	arrOfA1 := []A{{1}, {2}, {3}}
	arrOfA2 := []A{}
	arrOfB1 := []B{{1}, {2}, {3}, {4}}
	arrOfB2 := []B{}
	arrOfB3 := []B{{2}, {3}}

	expectSubsetA1B1 := true
	expectSubsetA1B2 := false
	expectSubsetA2B1 := true
	expectSubsetA2B2 := true
	expectSubsetA1B3 := false

	equalAB := func(a *A, b *B) bool {
		return a.va == b.vb
	}

	assert.Equal(t, expectSubsetA1B1, Subset(arrOfA1, arrOfB1, equalAB))
	assert.Equal(t, expectSubsetA1B2, Subset(arrOfA1, arrOfB2, equalAB))
	assert.Equal(t, expectSubsetA2B1, Subset(arrOfA2, arrOfB1, equalAB))
	assert.Equal(t, expectSubsetA2B2, Subset(arrOfA2, arrOfB2, equalAB))
	assert.Equal(t, expectSubsetA1B3, Subset(arrOfA1, arrOfB3, equalAB))
}

func TestSuperset(t *testing.T) {
	arrOfA1 := []A{{1}, {2}, {3}, {4}}
	arrOfA2 := []A{{2}, {3}}
	arrOfA3 := []A{}
	arrOfB1 := []B{{1}, {2}, {3}}
	arrOfB2 := []B{}

	expectSupersetA1B1 := true
	expectSupersetA1B2 := true
	expectSupersetA2B1 := false
	expectSupersetA2B2 := true
	expectSupersetA3B1 := false
	expectSupersetA3B2 := true

	equalAB := func(a *A, b *B) bool {
		return a.va == b.vb
	}

	assert.Equal(t, expectSupersetA1B1, Superset(arrOfA1, arrOfB1, equalAB))
	assert.Equal(t, expectSupersetA1B2, Superset(arrOfA1, arrOfB2, equalAB))
	assert.Equal(t, expectSupersetA2B1, Superset(arrOfA2, arrOfB1, equalAB))
	assert.Equal(t, expectSupersetA2B2, Superset(arrOfA2, arrOfB2, equalAB))
	assert.Equal(t, expectSupersetA3B1, Superset(arrOfA3, arrOfB1, equalAB))
	assert.Equal(t, expectSupersetA3B2, Superset(arrOfA3, arrOfB2, equalAB))
}

func TestContains(t *testing.T) {
	A1 := A{1}
	arrOfB1 := []B{{1}, {2}, {3}, {4}}
	arrOfB2 := []B{{2}, {3}}
	arrOfB3 := []B{}

	expectContainsA1B1 := true
	expectContainsA1B2 := false
	expectContainsA1B3 := false

	equalAB := func(a *A, b *B) bool {
		return a.va == b.vb
	}

	assert.Equal(t, expectContainsA1B1, Contains(&A1, arrOfB1, equalAB))
	assert.Equal(t, expectContainsA1B2, Contains(&A1, arrOfB2, equalAB))
	assert.Equal(t, expectContainsA1B3, Contains(&A1, arrOfB3, equalAB))
}
