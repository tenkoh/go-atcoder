package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(b []byte) int {
	// bの部分文字列の種類を全探索する
	m := map[string]struct{}{}
	for l := 1; l <= len(b); l++ {
		for i := 0; i+l <= len(b); i++ {
			m[string(b[i:i+l])] = struct{}{}
		}
	}
	return len(m)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	fmt.Println(solve(sc.Bytes()))
}
