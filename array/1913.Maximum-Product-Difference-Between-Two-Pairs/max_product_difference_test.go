package main

import "testing"

func TestMaxProductDifference(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "test #1",
			args: []int{5,6,2,7,4},
			want: 34,
		},
		{
			name: "test #2",
			args: []int{4,2,5,9,7,4,8},
			want: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProductDifference1(tt.args); got != tt.want {
				t.Errorf("maxProductDifference1() = %d, want %d", got, tt.want)
			}

			if got := maxProductDifference2(tt.args); got != tt.want {
				t.Errorf("maxProductDifference2() = %d, want %d", got, tt.want)
			}
		})
	}
}
