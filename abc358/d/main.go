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
	m := scan1(sc)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = scan1(sc)
	}
	bs := make([]int, m)
	for i := 0; i < m; i++ {
		bs[i] = scan1(sc)
	}

	// as, bsを昇順にソート
	// bsを順にたどり、asの中でbs[i]以上の最小のインデックスと値を求める
	sort.Ints(as)
	sort.Ints(bs)

	ans := 0
	for _, b := range bs {
		idx := sort.SearchInts(as, b)
		// idxがnの場合、asの中でb以上の値が存在しない
		if idx == len(as) {
			return -1
		}
		ans += as[idx]
		as = as[idx+1:]
	}
	return ans
}

func main() {
	fmt.Println(solve(os.Stdin))
}
