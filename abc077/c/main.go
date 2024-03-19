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

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	as := make([]int, n)
	bs := make([]int, n)
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = scan1(sc)
	}
	sort.Ints(as)
	for i := 0; i < n; i++ {
		bs[i] = scan1(sc)
	}
	sort.Ints(bs)
	for i := 0; i < n; i++ {
		cs[i] = scan1(sc)
	}
	sort.Ints(cs)

	// bを固定して a < b < c となるa, cの個数を求める
	count := 0
	for _, b := range bs {
		// bより小さいaの数
		if b <= as[0] {
			continue
		}
		ai := sort.SearchInts(as, b)

		// bより大きいcの数
		if cs[n-1] <= b {
			continue
		}
		ci := sort.SearchInts(cs, b+1)
		count += ai * (n - ci)
	}
	return count
}

func main() {
	fmt.Println(solve(os.Stdin))
}
