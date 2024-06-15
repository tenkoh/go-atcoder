package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `3 4
0 2 10
`
	s2 := `3 3
1 4 7
`
	s3 := `10 50000
120190 165111 196897 456895 540000 552614 561627 743796 757613 991216
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
		{"s1", args{r: strings.NewReader(s1)}, []int{4, 8, 14}},
		{"s2", args{r: strings.NewReader(s2)}, []int{4, 7, 10}},
		{"s3", args{r: strings.NewReader(s3)}, []int{170190, 220190, 270190, 506895, 590000, 640000, 690000, 793796, 843796, 1041216}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
