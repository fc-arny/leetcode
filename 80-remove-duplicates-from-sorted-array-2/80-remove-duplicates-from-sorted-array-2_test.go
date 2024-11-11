package leetcode

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "remove single element",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3, 3},
			},
			want: 6,
		},
		{
			name: "remove more than one element",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			},
			want: 7,
		},
		{
			name: "remove duplicates in all groups",
			args: args{
				nums: []int{0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3},
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeDuplicates(tt.args.nums)
			if result != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", result, tt.want)
			}
		})
	}
}
