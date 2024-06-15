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

func sToBit(s string) int {
	bit := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'o' {
			bit |= 1 << i
		}
	}
	return bit
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	m := scan1(sc)
	ss := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		ss[i] = sToBit(sc.Text())
	}

	ans := 1 << 60
	for bit := 0; bit < 1<<n; bit++ {
		// bitを2進数表記したときに1が立っているss[i]の要素のORをとっていく
		sbit := 0
		count := 0
		for i := 0; i < n; i++ {
			if bit>>i&1 == 0 {
				continue
			}
			count++
			sbit |= ss[i]
		}
		if sbit == 1<<m-1 && count < ans {
			ans = count
		}
	}
	return ans
}

func main() {
	fmt.Println(solve(os.Stdin))
}
