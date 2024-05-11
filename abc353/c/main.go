package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const mod = 100000000

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
	for i := 0; i < n; i++ {
		as[i] = scan1(sc)
	}

	var s int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			s += (as[i] + as[j]) % mod
		}
	}
	return s
}

func main() {
	fmt.Println(solve(os.Stdin))
}
