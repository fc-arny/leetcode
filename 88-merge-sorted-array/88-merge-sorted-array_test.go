package leetcode

import (
	"reflect"
	"testing"
)


func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
    nums2 []int
    n     int
	}

	tests := []struct {
		name string
		args args
		want []int
  }{
		{
			name: "merge from leetcode",
			args: args{	
				nums1: []int{1, 2, 3, 0, 0, 0},
				m:     3,
				nums2: []int{2, 5, 6},
				n:     3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name: "merge when nums1 is empty",
			args: args{	
				nums1: []int{0,0},
				m:     0,
				nums2: []int{1,2},
				n:     2,
			},
			want: []int{1, 2},
		},
		{
			name: "merge whennums2 is empty",
			args: args{
				nums1: []int{1, 2, 3},
				m:     3,
				nums2: []int{},
				n:     0,
			},
			want: []int{1, 2, 3},
		},
	}


	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			
			res := tt.args.nums1
			merge(res, tt.args.m, tt.args.nums2, tt.args.n)
			
			if !reflect.DeepEqual(res, tt.want){
				t.Errorf("merge() = %v, want %v", res, tt.want)
			}
    })
	}
}

