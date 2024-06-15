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
	a := scan1(sc)
	nextPurchase := 0

	ts := make([]int, n)
	for i := 0; i < n; i++ {
		t := scan1(sc)
		if t > nextPurchase {
			nextPurchase = t
		}
		nextPurchase += a
		ts[i] = nextPurchase
	}
	return ts
}

func main() {
	ans := solve(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		fmt.Fprintln(bw, v)
	}
	bw.Flush()
}
