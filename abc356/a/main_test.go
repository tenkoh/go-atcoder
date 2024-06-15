package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"s1", args{reader: strings.NewReader("5 2 3")}, []int{1, 3, 2, 4, 5}},
		{"s2", args{reader: strings.NewReader("7 1 1")}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"s3", args{reader: strings.NewReader("10 1 10")}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
