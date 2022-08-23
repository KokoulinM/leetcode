package main

import "testing"

func Test_numOfStrings(t *testing.T) {
	type args struct {
		patterns []string
		word     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{
				patterns: []string{"a", "abc", "bc", "d"},
				word:     "abc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfStrings(tt.args.patterns, tt.args.word); got != tt.want {
				t.Errorf("numOfStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
