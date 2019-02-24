package ex21

import "testing"

var baddCases = []struct {
	a []int
	b []int
	c []int
}{
	{a: []int{1, 1}, b: []int{1, 1}, c: []int{1, 0, 0}},
	{a: []int{0, 1}, b: []int{1, 0}, c: []int{0, 1, 1}},
	{a: []int{0, 1}, b: []int{0, 1}, c: []int{0, 1, 0}},
	{a: []int{0, 1}, b: []int{0, 0}, c: []int{0, 0, 1}},
	{a: []int{0, 0}, b: []int{0, 0}, c: []int{0, 0, 0}},
}

func TestBAdd(t *testing.T) {
	for _, tc := range baddCases {
		c := BAdd(tc.a, tc.b)

		if len(c) != len(tc.a)+1 {
			t.Fatalf("unexpected length: got %d, want %d", len(c), len(tc.a)+1)
		}

		for n := range tc.c {
			if c[n] != tc.c[n] {
				t.Fatalf("unexpected value: got %v, want %v", c, tc.c)
			}
		}
	}
}
