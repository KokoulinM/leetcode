package main

import "testing"

func Test_sortSentence(t *testing.T) {
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
				s: "is2 sentence4 This1 a3",
			},
			want: "This is a sentence",
		},
		{
			name: "Test #2",
			args: args{
				s: "Myself2 Me1 I4 and3",
			},
			want: "Me Myself and I",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortSentence(tt.args.s); got != tt.want {
				t.Errorf("sortSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
