package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) bool {
	b := []byte(s)
	m := map[byte]int{}
	for _, v := range b {
		m[v]++
	}

	mb := map[int]int{}
	for _, v := range m {
		mb[v]++
	}

	for _, v := range mb {
		if v != 2 {
			return false
		}
	}
	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if solve(sc.Text()) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
