package ex21

// Sort take a slice of int and returns a new slice with the elements of the
// underlying array sorted in ascending numerical order.
func Sort(a []int) {
	for j := 1; j < len(a); j++ {
		key := a[j]

		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
}

// Desc take a slice of int and returns a new slice with the elements of the
// underlying array sorted in descending numerical order.
func Desc(a []int) {
	for j := len(a) - 2; j >= 0; j-- {
		key := a[j]

		i := j + 1
		for i < len(a) && a[i] > key {
			a[i-1] = a[i]
			i = i + 1
		}
		a[i-1] = key
	}
}
