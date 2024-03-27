package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `5
00011
3 9 2 6 4
`
	s2 := `4
1001
1 2 3 4
`
	s3 := `11
11111100111
512298012 821282085 543342199 868532399 690830957 973970164 928915367 954764623 923012648 540375785 925723427
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"ex1", args{strings.NewReader(s1)}, 7},
		{"ex2", args{strings.NewReader(s2)}, 0},
		{"ex3", args{strings.NewReader(s3)}, 2286846953},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
