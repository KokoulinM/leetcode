package main

import "testing"

func Test_reversePrefix(t *testing.T) {
	type args struct {
		word string
		ch   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test #1",
			args: args{
				word: "abcdefd",
				ch:   'd',
			},
			want: "dcbaefd",
		},
		{
			name: "Test #2",
			args: args{
				word: "xyxzxe",
				ch:   'z',
			},
			want: "zxyxxe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrefix(tt.args.word, tt.args.ch); got != tt.want {
				t.Errorf("reversePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
