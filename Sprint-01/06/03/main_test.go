package main

import (
	"testing"
)

func TestSliceCopy(t *testing.T) {
	type test struct {
		numsData []int
		numsCap  int
		expected []int
	}

	tests := []test{
		{
			numsData: []int{1, 2, 3},
			numsCap:  10,
			expected: []int{1, 2, 3},
		},
		{
			numsData: []int{},
			numsCap:  10,
			expected: []int{},
		},
	}
	for _, tc := range tests {
		originNums := make([]int, len(tc.numsData), tc.numsCap)

		res := SliceCopy(originNums)

		if len(res) != len(originNums) {
			t.Fatalf("expected len: %v, got len: %v", len(tc.expected), len(res))
		}

		if cap(res) != len(originNums) {
			t.Fatalf("expected cap: %v, got cap: %v", len(tc.expected), cap(res))
		}

	}
}
