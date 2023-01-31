package structsetop

func Union() {

}

// Intersection intersects A with B by comparing each element with the given equal func
func Intersection[T any, U any, F ~func(*T, *U) bool](
	A []T,
	B []U,
	equal F,
) []T {
	result := []T{}

	if len(A) == 0 || len(B) == 0 {
		return result
	}

	for _, a := range A {
		for _, b := range B {
			if equal(&a, &b) {
				result = append(result, a)
				break
			}
		}
	}

	return result
}

func Difference[T any, U any, F ~func(*T, *U) bool](
	A []T,
	B []U,
	equal F,
) []T {

	result := []T{}

	if len(A) == 0 {
		return result
	}
	if len(B) == 0 {
		for _, a := range A {
			result = append(result, a)
		}
		return result
	}

	for _, a := range A {
		var inSetB bool

		for _, b := range B {
			if equal(&a, &b) {
				inSetB = true
				break
			}
		}

		if !inSetB {
			result = append(result, a)
		}
	}

	return result
}

func SymmetricDifference[T any, U any, F ~func(*T, *U) bool](
	A []T,
	B []U,
	equal F,
) ([]T, []U) {

	resultA := []T{}
	resultB := []U{}

	if len(A) == 0 && len(B) == 0 {
		return resultA, resultB
	}
	if len(A) == 0 {
		for _, b := range B {
			resultB = append(resultB, b)
		}
		return resultA, resultB
	}
	if len(B) == 0 {
		for _, a := range A {
			resultA = append(resultA, a)
		}
		return resultA, resultB
	}

	matchMapB := make(map[int]struct{})

	for _, a := range A {
		var inSetB bool

		for idx, b := range B {
			if equal(&a, &b) {
				inSetB = true
				matchMapB[idx] = struct{}{}
				break
			}

		}
		if !inSetB {
			resultA = append(resultA, a)
		}
	}

	for idx, b := range B {
		if _, ok := matchMapB[idx]; !ok {
			resultB = append(resultB, b)
		}
	}

	return resultA, resultB
}

func Subset() {

}

func Superset() {

}
