package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{s: "commencement"}, true},
		{"2", args{s: "banana"}, false},
		{"3", args{s: "ab"}, true},
		{"4", args{s: "a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
