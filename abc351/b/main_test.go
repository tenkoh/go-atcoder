package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_solve(t *testing.T) {
	s1 := `3
abc
def
ghi
abc
bef
ghi
`
	s2 := `1
f
q
`
	s3 := `10
eixfumagit
vtophbepfe
pxbfgsqcug
ugpugtsxzq
bvfhxyehfk
uqyfwtmglr
jaitenfqiq
acwvufpfvv
jhaddglpva
aacxsyqvoj
eixfumagit
vtophbepfe
pxbfgsqcug
ugpugtsxzq
bvfhxyehok
uqyfwtmglr
jaitenfqiq
acwvufpfvv
jhaddglpva
aacxsyqvoj
`
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want coord
	}{
		// TODO: Add test cases.
		{"1", args{r: strings.NewReader(s1)}, coord{2, 1}},
		{"2", args{r: strings.NewReader(s2)}, coord{1, 1}},
		{"3", args{r: strings.NewReader(s3)}, coord{5, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
