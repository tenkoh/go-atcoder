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

	var pos int
	for _, v := range bs {
		if v == bt[pos] {
			pos++
			if pos == 2 && bt[2] == 'x' {
				return true
			}
			if pos == 3 {
				return true
			}
		}
	}
	return false
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
