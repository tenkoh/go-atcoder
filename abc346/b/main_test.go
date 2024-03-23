package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		w int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"ex1", args{3, 2}, true},
		{"ex2", args{3, 0}, false},
		{"ex3", args{92, 66}, true},
		{"ex4", args{1, 0}, true},
		{"ex5", args{0, 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.w, tt.args.b); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
