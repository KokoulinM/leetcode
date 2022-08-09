package main

import (
	"reflect"
	"testing"
)

func Test_matrixHorizontal(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test #1",
			args: args{
				mat: [][]int{{1,2,3}, {4,5,6}, {7,8,9}},
			},
			want: []int{1,2,3,6,5,4,7,8,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixHorizontal(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixHorizontal() = %v, want %v", got, tt.want)
			}
		})
	}
}
