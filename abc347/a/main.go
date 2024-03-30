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
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		a := scan1(sc)
		if a%k == 0 {
			ans = append(ans, a/k)
		}
	}
	return ans
}

func main() {
	ans := solve(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	for i, a := range ans {
		if i > 0 {
			bw.WriteByte(' ')
		}
		fmt.Fprint(bw, a)
	}
	fmt.Fprintln(bw)
	bw.Flush()
}
