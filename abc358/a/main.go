package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	s := sc.Text()
	sc.Scan()
	t := sc.Text()

	if s == "AtCoder" && t == "Land" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
