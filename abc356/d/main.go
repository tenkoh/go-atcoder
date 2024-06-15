package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOD = 998244353

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func solve(N, M int) int {
	var result int = 0

	// 各ビット位置に対して計算
	for bit := 0; bit < 60; bit++ {
		// bit番目のビットが立っているkの個数を計算
		// 1 << (bit + 1)のサイクルで0,1が切り替わり、各サイクルに1<<bit個の1がある。
		bitCount := (N + 1) / (1 << (bit + 1)) * (1 << bit)
		remainder := (N + 1) % (1 << (bit + 1))
		if remainder > (1 << bit) {
			bitCount += remainder - (1 << bit)
		}

		if M&(1<<bit) != 0 {
			result += bitCount
			result %= MOD
		}
	}

	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	m := scan1(sc)
	fmt.Println(solve(n, m))
}
