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

func solve(r io.Reader) bool {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	m := scan1(sc)
	as := make([]int, m)
	for j := 0; j < m; j++ {
		as[j] = scan1(sc)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			as[j] -= scan1(sc)
		}
	}
	for _, a := range as {
		if a > 0 {
			return false
		}
	}
	return true
}

func main() {
	if solve(os.Stdin) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
