package ex23

import "math"

// Sort take a slice of int and returns a new slice with the elements of the
// underlying array sorted in ascending numerical order. Sort is performed
// using the merge-sort algorithm.
func Sort(A []int) {
	if len(A) <= 1 {
		return
	}

	q := len(A) / 2

	Sort(A[:q])
	Sort(A[q:])
	merge(A, q)
}

func merge(A []int, q int) {
	n1 := len(A[:q])
	n2 := len(A[q:])

	B := make([]int, n1+1)
	C := make([]int, n2+1)

	copy(B, A[:q])
	copy(C, A[q:])

	// We set an additional final element in each sub-array to be the
	// maximum allowed value to ensure we don't compare with an element
	// that exists outside the array bounds.
	B[n1] = math.MaxInt64
	C[n2] = math.MaxInt64

	i, j := 0, 0
	for k := range A {
		if B[i] <= C[j] {
			A[k] = B[i]
			i++
		} else {
			A[k] = C[j]
			j++
		}
	}
}
