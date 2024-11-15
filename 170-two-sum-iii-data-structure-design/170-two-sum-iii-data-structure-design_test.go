package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	type args struct {
		numbers []int
		value   int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "single element in array",
			args: args{numbers: []int{1}, value: 1},
			want: true,
		},
		{
			name: "no match in array",
			args: args{numbers: []int{1, 2, 3}, value: 4},
			want: true,
		},
		{
			name: "multiple matches in array",
			args: args{numbers: []int{1, 2, 3, 4, 4, 5}, value: 4},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsum := Constructor()
			for _, num := range tt.args.numbers {
				tsum.Add(num)
			}

			if got := tsum.Find(tt.args.value); got != tt.want {
				t.Errorf("TwoSum.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
