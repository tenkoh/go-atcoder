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

func solve(r io.Reader) bool {
	sc := bufio.NewScanner(bufio.NewReader(r))
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	m := scan1(sc)
	as := make([]int, n)
	bs := make([]int, m)
	for i := 0; i < n; i++ {
		as[i] = scan1(sc)
	}
	for i := 0; i < m; i++ {
		bs[i] = scan1(sc)
	}
	// as, bsをそれぞれ昇順に並び替えたあと、互いに先頭から比較して一番小さいものを取り出していく。
	// このとき、asから連続で取り出したらtrueを返す
	sort.Ints(as)
	sort.Ints(bs)

	tookAsHistory := make([]bool, 0, n+m)
	for i := 0; i < n+m; i++ {
		if len(as) == 0 {
			// asがからならbsから取り出す
			bs = bs[1:]
			tookAsHistory = append(tookAsHistory, false)
		} else if len(bs) == 0 {
			// bsがからならasから取り出す
			as = as[1:]
			tookAsHistory = append(tookAsHistory, true)
		} else {
			if as[0] < bs[0] {
				// asから取り出す
				as = as[1:]
				tookAsHistory = append(tookAsHistory, true)
			} else {
				// bsから取り出す
				bs = bs[1:]
				tookAsHistory = append(tookAsHistory, false)
			}
		}

		if i > 0 && tookAsHistory[i] && tookAsHistory[i-1] {
			return true
		}
	}
	return false
}

func main() {
	ans := solve(os.Stdin)
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
