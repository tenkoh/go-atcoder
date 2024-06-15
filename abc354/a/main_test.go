package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		h int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{54}, 6},
		{"2", args{7}, 4},
		{"3", args{262144}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.h); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
