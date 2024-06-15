package main

import (
	"io"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `3
takahashi 2
aoki 6
snuke 5
`
	s2 := `3
takahashi 2813
takahashixx 1086
takahashix 4229
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{r: io.Reader(strings.NewReader(s1))}, "snuke"},
		{"2", args{r: io.Reader(strings.NewReader(s2))}, "takahashix"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
