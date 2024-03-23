package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4 5
1 6 3 1
`
	s2 := `1 3
346
`
	s3 := `10 158260522
877914575 24979445 623690081 262703497 24979445 1822804784 1430302156 1161735902 923078537 1189330739
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"ex1", args{strings.NewReader(s1)}, 11},
		{"ex2", args{strings.NewReader(s2)}, 6},
		{"ex3", args{strings.NewReader(s3)}, 12523196466007058},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
