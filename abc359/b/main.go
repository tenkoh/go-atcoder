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
	// i:=0, 1, ..., 2*(n-1)-1まで、a[i]==a[i+2]となる数を数える
	as := make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		as[i] = scan1(sc)
	}
	var count int
	for i := 0; i < 2*(n-1); i++ {
		if as[i] == as[i+2] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(solve(os.Stdin))
}
