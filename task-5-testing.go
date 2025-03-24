package main

import "testing"

func TestExistIntersect(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected bool
	}{
		{a: []int{1, 2, 3}, b: []int{4, 5, 6}, expected: false},
		{a: []int{65, 3, 58, 678, 64}, b: []int{64, 2, 3, 43}, expected: true},
	}

	for _, test := range tests {
		result := existIntersect(test.a, test.b)
		if result != test.expected {
			t.Fatalf("For input a: %v and b: %v, expected %v but got %v", test.a, test.b, test.expected, result)
		}
	}
}
