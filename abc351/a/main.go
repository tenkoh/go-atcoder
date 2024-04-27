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

	var ta, tb int
	for i := 0; i < 9; i++ {
		ta += scan1(sc)
	}
	for i := 0; i < 8; i++ {
		tb += scan1(sc)
	}
	return ta - tb + 1
}

func main() {
	fmt.Println(solve(os.Stdin))
}
