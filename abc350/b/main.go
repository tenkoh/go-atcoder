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
	q := scan1(sc)
	teeth := map[int]bool{}
	for i := 0; i < n; i++ {
		// 1-indexed
		teeth[i+1] = true
	}
	for i := 0; i < q; i++ {
		t := scan1(sc)
		teeth[t] = !teeth[t]
	}
	ans := 0
	for i := 1; i <= n; i++ {
		if teeth[i] {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(solve(os.Stdin))
}
