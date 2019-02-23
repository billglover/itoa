package ex21

// BAdd takes two n-element arrays, representing an n-bit binary numbers. It
// returns an (n+1)-element array representing an (n+1)-bit binary number that
// is the sum of the inputs.
func BAdd(a, b []int) []int {
	if len(a) != len(b) {
		return nil
	}

	c := make([]int, len(a)+1)
	carry := 0

	for i := len(a) - 1; i >= 0; i-- {
		c[i+1] = a[i] + b[i] + carry
		if c[i+1] > 1 {
			c[i+1] = 0
			carry = 1
		} else {
			carry = 0
		}
	}
	c[0] = carry

	return c
}
