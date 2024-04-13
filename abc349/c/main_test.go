package main

import (
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{s: "narita", t: "NRT"}, true},
		{"2", args{s: "losangeles", t: "LAX"}, true},
		{"3", args{s: "snuke", t: "RNG"}, false},
		{"4", args{s: "aoyama", t: "AAA"}, true},
		{"too long", args{s: strings.Repeat("nagoya", 10000), t: "NGY"}, true},
		{"wrong order", args{s: "nagoya", t: "YGN"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
