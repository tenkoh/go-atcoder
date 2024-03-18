package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	sum := 0
	ps := make([]int, n)
	for i := 0; i < n; i++ {
		ps[i] = scan1(sc)
		sum += ps[i]
	}
	sort.Ints(ps)

	dp := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]bool, sum+1)
	}
	dp[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j < sum; j++ {
			if !dp[i][j] {
				continue
			}
			// 足さない時
			dp[i+1][j] = true
			// 足す時
			dp[i+1][j+ps[i]] = true
		}
	}

	count := 0
	// sumまで数え上げて取りうる数を数える
	for j := 0; j < sum+1; j++ {
		for i := 0; i < n+1; i++ {
			if dp[i][j] {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(solve(os.Stdin))
}
