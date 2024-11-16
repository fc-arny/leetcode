package leetcode

import "testing"

func TestTitleToUumber(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single digit",
			input: "A",
			want:  1,
		},
		{
			name:  "multiple digits",
			input: "AB",
			want:  28,
		},
		{
			name:  "capital letters",
			input: "ZY",
			want:  701,
		},
		// Add more test cases as needed...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.input); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
