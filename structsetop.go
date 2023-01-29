package structsetop

func Union() {

}

// Intersection intersects A with B by comparing each element with the given compare func
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

func Difference() {

}

func SymmetricDifference() {

}

func Subset() {

}

func Superset() {

}
