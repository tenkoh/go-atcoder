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

func solve(reader io.Reader) []int {
	sc := bufio.NewScanner(reader)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	l := scan1(sc)
	r := scan1(sc)

	ans := make([]int, n)
	for i := 1; i <= n; i++ {
		if i >= l && i <= r {
			ans[i-1] = r - (i - l)
		} else {
			ans[i-1] = i
		}
	}
	return ans
}

func main() {
	ans := solve(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	for i, v := range ans {
		if i > 0 {
			bw.WriteByte(' ')
		}
		fmt.Fprint(bw, v)
	}
	fmt.Fprint(bw, "\n")
	bw.Flush()
}
