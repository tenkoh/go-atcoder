package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(h int) int {
	var i, acc int
	for {
		acc += 1 << i
		if acc > h {
			return i + 1
		}
		i++
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	h, _ := strconv.Atoi(sc.Text())
	fmt.Println(solve(h))
}
