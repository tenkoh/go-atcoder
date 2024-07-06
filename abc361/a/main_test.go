package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `4 3 7
2 3 5 11
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"s1", args{strings.NewReader(s1)}, []int{2, 3, 5, 7, 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
