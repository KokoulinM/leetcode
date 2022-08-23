package main

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode {
						Val: 2,
						Left: nil,
						Right: nil,
					},
					Right: &TreeNode {
						Val: 3,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode {
								Val: 2,
								Left: nil,
								Right: nil,
							},
							Right: nil,
						},
					},
				},
			},
			want: 4,
		},
		{
			name: "Test #2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode {
						Val: 2,
						Left: nil,
						Right: nil,
					},
					Right: &TreeNode {
						Val: 3,
						Left: &TreeNode{
							Val: 5,
							Left: nil,
							Right: nil,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
