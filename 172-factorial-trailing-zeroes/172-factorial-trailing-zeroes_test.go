package leetcode

import "testing"

func TestTraillingZeroes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "n=3, n!=6",
			n:    3,
			want: 0,
		},
		{
			name: "n=5, n!=120",
			n:    5,
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
