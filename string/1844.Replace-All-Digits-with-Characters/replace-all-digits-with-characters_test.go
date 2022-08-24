package main

import "testing"

func Test_replaceDigits(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test #1",
			args: args{
				s: "a1c1e1",
			},
			want: "abcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceDigits(tt.args.s); got != tt.want {
				t.Errorf("replaceDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
