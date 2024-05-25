package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func solve(a, b int) int {
	if a == b {
		return -1
	}
	ans := make([]bool, 3)
	ans[a-1] = true
	ans[b-1] = true
	for i, v := range ans {
		if !v {
			return i + 1
		}
	}
	return -1
}

func main() {
	sc := bufio.NewScanner(bufio.NewReader(os.Stdin))
	sc.Split(bufio.ScanWords)
	a := scan1(sc)
	b := scan1(sc)
	fmt.Println(solve(a, b))
}
