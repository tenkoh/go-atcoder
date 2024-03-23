package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"ex1", args{strings.NewReader("3\n3 4 6\n")}, []int{12, 24}},
		{"ex2", args{strings.NewReader("5\n22 75 26 45 72\n")}, []int{1650, 1950, 1170, 3240}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
