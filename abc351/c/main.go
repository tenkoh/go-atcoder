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

	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		stack = append(stack, scan1(sc))
		for {
			l := len(stack)

			if l <= 1 {
				break
			}

			if stack[l-1] != stack[l-2] {
				break
			}

			stack[l-2] += 1
			stack = stack[:l-1]
		}
	}
	return len(stack)
}

func main() {
	fmt.Println(solve(os.Stdin))
}
