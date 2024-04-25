package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type pair struct {
	a, b int
}

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func solve(r io.Reader) []*pair {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)

	as := make([]int, n)
	pos := map[int]int{} // value=iのindex. 1-indexed
	for i := 0; i < n; i++ {
		as[i] = scan1(sc)
		pos[as[i]] = i + 1
	}

	var pairs []*pair
	// 1,2...nに並び替え
	for i := 1; i <= n; i++ {
		if i == pos[i] {
			continue
		}
		// 1-indexed
		pairs = append(pairs, &pair{i, pos[i]})
		as[i-1], as[pos[i]-1], pos[i], pos[as[i-1]] = as[pos[i]-1], as[i-1], pos[as[i-1]], pos[i]
	}
	return pairs
}

func main() {
	pairs := solve(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	fmt.Fprintln(bw, len(pairs))
	for _, p := range pairs {
		fmt.Fprintln(bw, p.a, p.b)
	}
}
