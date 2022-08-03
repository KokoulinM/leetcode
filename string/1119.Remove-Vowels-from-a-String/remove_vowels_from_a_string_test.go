package main

import (
	"testing"
)

func Test_removeVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test #1",
			args: args{
				s: "leetcodeisacommunityforcoders",
			},
			want: "ltcdscmmntyfrcdrs",
		},
		{
			name: "test #2",
			args: args{
				s: "aeiou",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeVowels(tt.args.s); got != tt.want {
				t.Errorf("removeVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
