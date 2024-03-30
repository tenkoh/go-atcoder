package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `3 2 5
1 2 9
`
	s2 := `2 5 10
10 15
`
	s3 := `4 347 347
347 700 705 710
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{r: strings.NewReader(s1)}, true},
		{"2", args{r: strings.NewReader(s2)}, false},
		{"3", args{r: strings.NewReader(s3)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
