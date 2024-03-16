package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func scan3(sc *bufio.Scanner) (a, b, c int) {
	sc.Scan()
	a, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	c, _ = strconv.Atoi(sc.Text())
	return
}

func chmax[T int | int64 | float32 | float64](a *T, b T) {
	if *a < b {
		*a = b
	}
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	as := make([]int, n)
	bs := make([]int, n)
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		as[i], bs[i], cs[i] = scan3(sc)
	}

	dp := make([][3]int, n)
	for i := 0; i < n; i++ {
		dp[i] = [3]int{0, 0, 0}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = as[i]
			dp[i][1] = bs[i]
			dp[i][2] = cs[i]
		} else {
			chmax(&dp[i][0], dp[i-1][1]+as[i])
			chmax(&dp[i][0], dp[i-1][2]+as[i])

			chmax(&dp[i][1], dp[i-1][0]+bs[i])
			chmax(&dp[i][1], dp[i-1][2]+bs[i])

			chmax(&dp[i][2], dp[i-1][0]+cs[i])
			chmax(&dp[i][2], dp[i-1][1]+cs[i])
		}
	}

	ans := 0
	chmax(&ans, dp[n-1][0])
	chmax(&ans, dp[n-1][1])
	chmax(&ans, dp[n-1][2])
	return ans
}

func main() {
	fmt.Println(solve(os.Stdin))
}
