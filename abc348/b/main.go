package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p *point) distance(q *point) int {
	return (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)
}

func scan2(sc *bufio.Scanner) *point {
	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	return &point{x, y}
}

func solve(r io.Reader) []int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	points := make([]*point, n)
	for i := 0; i < n; i++ {
		points[i] = scan2(sc)
	}

	distances := make([]int, n)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		p := points[i]
		for j := 0; j < n; j++ {
			d := p.distance(points[j])
			if d == distances[i] {
				if j < ans[i] {
					ans[i] = j
				}
				continue
			}
			if d > distances[i] {
				distances[i] = d
				ans[i] = j
			}
		}
	}

	// 点の番号は1-indexed
	for i := 0; i < n; i++ {
		ans[i]++
	}
	return ans
}

func main() {
	ans := solve(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	for _, v := range ans {
		fmt.Fprintln(bw, v)
	}
	bw.Flush()
}
