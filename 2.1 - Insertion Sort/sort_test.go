package ex21

import "testing"

var sortCases = []struct {
	i []int
	o []int
}{
	{i: []int{1, 2, 3, 4, 5}, o: []int{1, 2, 3, 4, 5}},
	{i: []int{5, 4, 3, 2, 1}, o: []int{1, 2, 3, 4, 5}},
	{i: []int{5, 2, 4, 6, 1, 3}, o: []int{1, 2, 3, 4, 5, 6}},
	{i: []int{31, 41, 59, 26, 41, 58}, o: []int{26, 31, 41, 41, 58, 59}},
}

func TestSort(t *testing.T) {
	for _, c := range sortCases {
		s := Sort(c.i)

		if len(s) != len(c.o) {
			t.Fatalf("unexpected length: got %d, want %d", len(s), len(c.o))
		}

		for n := range c.o {
			if s[n] != c.o[n] {
				t.Fatalf("unexpected value: got %v, want %v", s, c.o)
			}
		}
	}
}

var descCases = []struct {
	i []int
	o []int
}{
	{i: []int{5, 4, 3, 2, 1}, o: []int{5, 4, 3, 2, 1}},
	{i: []int{1, 2, 3, 4, 5}, o: []int{5, 4, 3, 2, 1}},
	{i: []int{5, 2, 4, 6, 1, 3}, o: []int{6, 5, 4, 3, 2, 1}},
	{i: []int{31, 41, 59, 26, 41, 58}, o: []int{59, 58, 41, 41, 31, 26}},
}

func TestDesc(t *testing.T) {
	for _, c := range descCases {
		s := Desc(c.i)

		if len(s) != len(c.o) {
			t.Fatalf("unexpected length: got %d, want %d", len(s), len(c.o))
		}

		for n := range c.o {
			if s[n] != c.o[n] {
				t.Fatalf("unexpected value: got %v, want %v", s, c.o)
			}
		}
	}
}
