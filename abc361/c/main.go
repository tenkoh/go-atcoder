package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func chmin[T int | int64 | float32 | float64](a *T, b T) {
	if *a > b {
		*a = b
	}
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	k := scan1(sc)
	as := make([]int, n)
	for i := range as {
		as[i] = scan1(sc)
	}

	if n == k-1 {
		return 0
	}

	sort.Ints(as)
	diffmin := 1 << 60
	for leftTrim := 0; leftTrim <= k; leftTrim++ {
		rightTrim := k - leftTrim
		diff := as[n-rightTrim-1] - as[leftTrim]
		chmin(&diffmin, diff)
	}
	return diffmin
}

func main() {
	fmt.Println(solve(os.Stdin))
}
