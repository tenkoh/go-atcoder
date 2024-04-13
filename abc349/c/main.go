package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(s, t string) bool {
	bs := []byte(s)
	bt := []byte(strings.ToLower(t))

	mbs := map[byte]int{}
	for _, v := range bs {
		mbs[v]++
	}
	for i, v := range bt {
		if i == 2 && v == 'x' {
			return true
		}
		if mbs[v] == 0 {
			return false
		}
		mbs[v]--
	}
	return true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 200000), 0)
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()
	if solve(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
