package main

import "testing"

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//{
		//	name: "Test #1",
		//	args: args{
		//		paths: [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
		//	},
		//	want: "Sao Paulo",
		//},
		{
			name: "Test #2",
			args: args{
				paths: [][]string{{"D", "B"}, {"C", "A"}, {"B", "C"}},
			},
			want: "A",
		},
		//{
		//	name: "Test #3",
		//	args: args{
		//		paths: [][]string{{"A", "Z"}},
		//	},
		//	want: "Z",
		//},
		//{
		//	name: "Test #4",
		//	args: args{
		//		paths: [][]string{{"pYyNGfBYbm", "wxAscRuzOl"}, {"kzwEQHfwce", "pYyNGfBYbm"}},
		//	},
		//	want: "wxAscRuzOl",
		//},
		//{
		//	name: "Test #5",
		//	args: args{
		//		paths: [][]string{{"kzwEQHfwce", "pYyNGfBYbm"}, {"pYyNGfBYbm", "wxAscRuzOl"}},
		//	},
		//	want: "wxAscRuzOl",
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := destCity(tt.args.paths); got != tt.want {
				t.Errorf("destCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
