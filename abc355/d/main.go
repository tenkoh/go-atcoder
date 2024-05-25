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

	ls := make([]int, n)
	rs := make([]int, n)

	for i := 0; i < n; i++ {
		ls[i] = scan1(sc)
		rs[i] = scan1(sc)
	}
	sort.Ints(ls)
	sort.Ints(rs)

	// lsのそれぞれの要素に対して r < l となるrsの要素数を数える
	// 数え上げには尺取り法を使う
	// 組み合わせの総数から数え上げた数を引く
	ans := n * (n - 1) / 2
	j := 0
	for _, l := range ls {
		for rs[j] < l {
			j++
		}
		ans -= j
	}
	return ans
}

func main() {
	fmt.Println(solve(os.Stdin))
}
