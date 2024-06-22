package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		sx int
		sy int
		tx int
		ty int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"s1", args{5, 0, 2, 5}, 5},
		{"s2", args{3, 1, 4, 1}, 0},
		{"s3", args{2552608206527595, 5411232866732612, 771856005518028, 7206210729152763}, 1794977862420151},
		{"s4", args{0, 0, 1, 0}, 0},
		{"s4", args{0, 0, 2, 0}, 1},
		{"s5", args{0, 1, 0, 2}, 1},
		{"s6", args{0, 1, 6, 2}, 4},
		{"s6", args{0, 1, 7, 2}, 4},
		{"s7", args{0, 1, -2, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.sx, tt.args.sy, tt.args.tx, tt.args.ty); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
