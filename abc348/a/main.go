package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(n int) string {
	ans := make([]byte, n)
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 {
			ans[i] = 'x'
		} else {
			ans[i] = 'o'
		}
	}
	return string(ans)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	fmt.Println(solve(n))
}
