package main

import (
	"reflect"
	"testing"
)

func TestClean(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Ok!",
			args: args{
				nums: []int{6, 5, 4, 6, 2, 4, 6, 2, 1, 6},
				x:    6,
			},
			want: []int{5, 4, 2, 4, 2, 1},
		},
		{
			name: "Ok!",
			args: args{
				nums: []int{-1, 6, 4, 5, 1, 6, 211, 90},
				x:    0,
			},
			want: []int{-1, 6, 4, 5, 1, 6, 211, 90},
		},
		{
			name: "Ok!",
			args: args{
				nums: []int{},
				x:    1,
			},
			want: []int{},
		},
		{
			name: "Ok!",
			args: args{
				nums: []int{5, 5},
				x:    5,
			},
			want: []int{},
		},
		{
			name: "Ok!",
			args: args{
				nums: []int{3, -7},
				x:    -7,
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clean(tt.args.nums, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clean() = %v, want %v", got, tt.want)
			}
		})
	}
}
