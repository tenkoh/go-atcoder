package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	var count int
	for sc.Scan() {
		if sc.Text() == "Takahashi" {
			count++
		}
	}
	fmt.Println(count)
}
