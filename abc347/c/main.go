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

func pmod(a, b int) int {
	return (a%b + b) % b
}

func solve(r io.Reader) bool {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	a := scan1(sc)
	b := scan1(sc)
	wds := make([]int, 0, n+1)
	mwds := map[int]struct{}{}
	for i := 0; i < n; i++ {
		wd := scan1(sc) % (a + b)
		if _, ok := mwds[wd]; ok {
			continue
		}
		mwds[wd] = struct{}{}
		wds = append(wds, wd)
	}
	// 1日しか予定がない場合は無条件でYes
	if len(wds) == 1 {
		return true
	}
	sort.Ints(wds)
	wds = append(wds, wds[0])
	for i := 0; i < len(wds)-1; i++ {
		if pmod(wds[i+1]-wds[i], a+b) > b {
			return true
		}
	}
	return false
}

func main() {
	if solve(os.Stdin) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
