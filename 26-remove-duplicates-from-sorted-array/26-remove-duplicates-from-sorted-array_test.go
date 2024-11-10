package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "empty array",
			args: []int{},
			want: 0,
		},
		{
			name: "array with no duplicates",
			args: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "array with duplicates",
			args: []int{1, 1, 2, 2, 3, 4, 5, 5},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
