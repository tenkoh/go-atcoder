package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func solve(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)

	var totalRate int
	var users []string
	for i := 0; i < n; i++ {
		sc.Scan()
		users = append(users, sc.Text())
		totalRate += scan1(sc)
	}

	sort.Strings(users)
	return users[totalRate%n]
}

func main() {
	fmt.Println(solve(os.Stdin))
}
