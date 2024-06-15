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

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	k := scan1(sc)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = scan1(sc)
	}

	var count int
	tk := k
	for len(as) > 0 {
		if tk < as[0] {
			count++
			tk = k
			continue
		}
		tk -= as[0]
		as = as[1:]
	}
	count++
	return count
}

func main() {
	fmt.Println(solve(os.Stdin))
}
