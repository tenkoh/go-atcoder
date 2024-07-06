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
	k := scan1(sc)
	x := scan1(sc)
	as := make([]int, n)
	for i := range as {
		as[i] = scan1(sc)
	}

	ans := make([]int, n+1)
	copy(ans, as[:k])
	copy(ans[k:], []int{x})
	copy(ans[k+1:], as[k:])
	return ans
}

func main() {
	ans := solve(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	for i, a := range ans {
		if i != 0 {
			bw.WriteByte(' ')
		}
		fmt.Fprintf(bw, "%d", a)
	}
	bw.Flush()
}
