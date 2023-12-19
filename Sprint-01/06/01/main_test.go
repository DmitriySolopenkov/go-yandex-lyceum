package main

import (
	"reflect"
	"testing"
)

func TestUnderLimit(t *testing.T) {
	type test struct {
		nums      []int
		n         int
		limit     int
		expected  []int
		wantError bool
	}

	tests := []test{
		{
			nums:      []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12},
			n:         5,
			limit:     3,
			expected:  []int{2, 0, 2},
			wantError: false,
		},
		{
			nums:      nil,
			wantError: true,
		},
		{
			nums:      []int{},
			n:         5,
			limit:     3,
			expected:  []int{},
			wantError: false,
		},
		{
			nums:      []int{3, 5, 6},
			n:         5,
			limit:     10,
			expected:  []int{3, 5, 6},
			wantError: false,
		},
		{
			nums:      []int{-13, 0, 6},
			n:         1,
			limit:     -5,
			expected:  []int{-13},
			wantError: false,
		},
		{
			nums:      []int{},
			n:         -1,
			limit:     5,
			wantError: true,
		},
	}

	for _, tc := range tests {
		result, err := UnderLimit(tc.nums, tc.limit, tc.n)
		if tc.wantError {
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
		} else {
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("expected %v, got %v", tc.expected, result)
			}
		}
	}
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := UnderLimit(tt.args.nums, tt.args.limit, tt.args.n)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("UnderLimit() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("UnderLimit() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
