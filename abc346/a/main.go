package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func solve(r io.Reader) []int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	n := scan1(sc)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = scan1(sc)
	}

	bs := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		bs[i] = as[i] * as[i+1]
	}
	return bs
}

func main() {
	bs := solve(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	for i, b := range bs {
		if i > 0 {
			bw.WriteString(" ")
		}
		fmt.Fprint(bw, b)
	}
	bw.Flush()
}
