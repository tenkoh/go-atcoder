package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(s string) bool {
	if s == "ABC316" {
		return false
	}
	if s == "ABC000" {
		return false
	}
	if s < "ABC350" {
		return true
	}
	return false
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
