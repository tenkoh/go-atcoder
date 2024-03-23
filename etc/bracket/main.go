package main

import (
	"bufio"
	"fmt"
	"io"
)

type bracket struct {
	pos   int
	right bool
}

func solve(r io.Reader) (bool, []string) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanBytes)

	var stack []bracket
	var pairs []string
	for i := 0; sc.Scan(); i++ {
		s := sc.Text()
		b := bracket{pos: i}
		if s == ")" {
			b.right = true
		}
		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}
		prev := stack[len(stack)-1]
		if b.right && !prev.right {
			pairs = append(pairs, fmt.Sprintf("%d %d", prev.pos, i))
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, b)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return len(stack) == 0, pairs
}
