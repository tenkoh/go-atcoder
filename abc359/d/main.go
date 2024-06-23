package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func reverse(s string) string {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[len(rs)-1-i] = rs[len(rs)-1-i], rs[i]
	}
	return string(rs)
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	sc.Scan() // skip n
	k := scan1(sc)
	sc.Scan()
	s := sc.Text()

	dp := map[string]int{strings.Repeat("C", k-1): 1}
	for _, b := range s {
		tmp := map[string]int{}
		if b != 'B' {
			for k, v := range dp {
				tmp[k+"A"] = v
			}
		}
		if b != 'A' {
			for k, v := range dp {
				tmp[k+"B"] = v
			}
		}

		// dpを初期化
		dp = map[string]int{}
		for k, v := range tmp {
			if k == reverse(k) {
				continue
			}
			// この時点で剰余を取らないとオーバーフローする
			dp[k[1:]] += v % 998244353
		}
	}

	var sum int
	for _, v := range dp {
		sum += v
		sum %= 998244353
	}
	return sum
}

func main() {
	fmt.Println(solve(os.Stdin))
}
