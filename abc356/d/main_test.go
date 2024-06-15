package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N int
		M int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"s1", args{N: 4, M: 3}, 4},
		{"s2", args{N: 0, M: 0}, 0},
		{"s3", args{N: 1152921504606846975, M: 1152921504606846975}, 499791890},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.M); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
