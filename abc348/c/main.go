package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func scan2(sc *bufio.Scanner) (int, int) {
	sc.Scan()
	ac := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(ac[0])
	c, _ := strconv.Atoi(ac[1])
	return a, c
}

func solve(r io.Reader) int {
	// Ai, Ci で与えられる組み合わせの中で
	// CiごとのAiの最小値を求める
	// その最小値の中で最大のものを求める
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	m := map[int]int{}
	for i := 0; i < n; i++ {
		a, c := scan2(sc)
		if v, ok := m[c]; ok {
			if a < v {
				m[c] = a
			}
		} else {
			m[c] = a
		}
	}

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(solve(os.Stdin))
}
