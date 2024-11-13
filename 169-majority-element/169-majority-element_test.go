package leetcode

import (
	"testing"
)

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LC tescase #1",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "LC testcase #2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
		{
			name: "LC testcase #3",
			args: args{
				nums: []int{1, 2, 1, 3, 1, 4, 1, 1, 5},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := majorityElement(tt.args.nums)
			if result != tt.want {
				t.Errorf("mamajorityElement() = %v, want %v", result, tt.want)
			}
		})
	}
}
