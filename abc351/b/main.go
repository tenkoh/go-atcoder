package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type coord struct {
	i, j int // 1-indexed
}

func solve(r io.Reader) coord {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	ma := make(map[coord]rune)

	for i := 1; i <= n; i++ {
		sc.Scan()
		s := sc.Text()
		for j, r := range s {
			ma[coord{i, j + 1}] = r
		}
	}
	for i := 1; i <= n; i++ {
		sc.Scan()
		s := sc.Text()
		for j, r := range s {
			if ma[coord{i, j + 1}] != r {
				return coord{i, j + 1}
			}
		}
	}
	return coord{}
}

func main() {
	c := solve(os.Stdin)
	fmt.Printf("%d %d\n", c.i, c.j)
}
