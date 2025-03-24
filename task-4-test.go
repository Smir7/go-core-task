package main

import "fmt"

func TestTwoSlice(t []string) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{
			slice1:   []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			slice2:   []string{"banana", "date", "fig"},
			expected: []string{"apple", "cherry", "43", "lead", "gno1"},
		},
	}

	for _, test := range tests {
		result := twoSlice(test.slice1, test.slice2)
		if !equalSlices(result, test.expected) {
			fmt.Errorf("for input %v and %v, expected %v but got %v", test.slice1, test.slice2, test.expected, result)
		}
	}
}

func equalSlices(sliceA, sliceB []string) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}
	for i, v := range sliceA {
		if v != sliceB[i] {
			return false
		}
	}
	return true
}
