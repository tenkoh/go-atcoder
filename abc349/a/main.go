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
	var sum int
	for i := 0; i < n-1; i++ {
		sum += scan1(sc)
	}
	return -sum
}

func main() {
	fmt.Println(solve(os.Stdin))
}
