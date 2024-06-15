package main

import (
	"io"
	"strings"
	"testing"
)

func Test_sToBit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"s1", args{"oxo"}, 0b101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sToBit(tt.args.s); got != tt.want {
				t.Errorf("sToBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve(t *testing.T) {
	s1 := `3 5
oooxx
xooox
xxooo
`
	s2 := `3 2
oo
ox
xo
`
	s3 := `8 6
xxoxxo
xxoxxx
xoxxxx
xxxoxx
xxoooo
xxxxox
xoxxox
oxoxxo
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
		{"s1", args{r: strings.NewReader(s1)}, 2},
		{"s2", args{r: strings.NewReader(s2)}, 1},
		{"s3", args{r: strings.NewReader(s3)}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
