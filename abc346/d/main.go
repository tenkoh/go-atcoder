package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const inf = 1 << 60

func chmin[T int | int64 | float32 | float64](a *T, b T) {
	if *a > b {
		*a = b
	}
}

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 400000), 0)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	sc.Scan()
	s := sc.Text()
	cs := make([]int, n)
	for i := 0; i < n; i++ {
		cs[i] = scan1(sc)
	}

	// dpの初期化
	// dp[i][j][k]: i文字目までにj組のペアがあり、i文字目がkの時の総コスト
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		// どうせ2組以上のペアは気にしない
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, 2)
			dp[i][j][0] = inf
			dp[i][j][1] = inf
		}
	}
	dp[0][0][s[0]-'0'] = 0
	dp[0][0][(s[0]-'0')^1] = cs[0]

	for i := 1; i < n; i++ {
		r := int(s[i] - '0')
		for k := 0; k < 2; k++ {
			dp[i][0][k] = dp[i-1][0][k^1]
			if r != k {
				dp[i][0][k] += cs[i]
			}

			chmin(&dp[i][1][k], dp[i-1][0][k])
			chmin(&dp[i][1][k], dp[i-1][1][k^1])
			if r != k {
				dp[i][1][k] += cs[i]
			}
		}
	}
	ans := inf
	chmin(&ans, dp[n-1][1][0])
	chmin(&ans, dp[n-1][1][1])
	return ans
}

func main() {
	fmt.Println(solve(os.Stdin))
}
