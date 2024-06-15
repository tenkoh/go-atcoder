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
	h1 := scan1(sc)
	for i := 2; i <= n; i++ {
		h := scan1(sc)
		if h > h1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(solve(os.Stdin))
}
