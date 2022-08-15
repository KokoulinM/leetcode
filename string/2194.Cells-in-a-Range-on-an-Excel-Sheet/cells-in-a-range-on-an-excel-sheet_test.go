package main

import (
	"reflect"
	"testing"
)

func Test_cellsInRange(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test #1",
			args: args{
				s: "A1:F1",
			},
			want: []string{"A1", "B1", "C1", "D1", "E1", "F1"},
		},
		{
			name: "Test #2",
			args: args{
				s: "K1:L2",
			},
			want: []string{"K1", "K2", "L1", "L2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cellsInRange(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cellsInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
