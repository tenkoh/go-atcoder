package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func scan1(sc *bufio.Scanner) int64 {
	sc.Scan()
	n, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return n
}

func solve(r io.Reader) int64 {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	k := scan1(sc)
	ma := map[int64]struct{}{}
	var i int64
	// 重複を除去した上で後から引くべき値を控えておく
	for i = 0; i < n; i++ {
		m := scan1(sc)
		if m > k {
			continue
		}
		ma[m] = struct{}{}
	}
	var ans int64
	ans = k * (k + 1) / 2
	for m := range ma {
		ans -= m
	}
	return ans
}

func main() {
	bw := bufio.NewWriter(os.Stdout)
	fmt.Fprint(bw, solve(os.Stdin))
	bw.Flush()
}
