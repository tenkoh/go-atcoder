package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{1, 2}, 3},
		{"2", args{2, 3}, 1},
		{"3", args{1, 1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
