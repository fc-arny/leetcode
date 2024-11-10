package leetcode

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "remove single element",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "remove multiple elements",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "no elements to remove",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
				val:  6,
			},
			want: 5,
		},
		{
			name: "empty array",
			args: args{
				nums: []int{},
				val:  1,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElement(tt.args.nums, tt.args.val)
			if result != tt.want {
				t.Errorf("removeElement() = %v, want %v", result, tt.want)
			}
		})
	}
}
