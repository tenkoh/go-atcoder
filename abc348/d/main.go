package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// 深さ優先探索でスタートからゴールに辿り着けるかを判定する

func scan2(sc *bufio.Scanner) (int, int) {
	sc.Scan()
	vs := strings.Split(sc.Text(), " ")
	v1, _ := strconv.Atoi(vs[0])
	v2, _ := strconv.Atoi(vs[1])
	return v1, v2
}

func scan3(sc *bufio.Scanner) (int, int, int) {
	sc.Scan()
	vs := strings.Split(sc.Text(), " ")
	v1, _ := strconv.Atoi(vs[0])
	v2, _ := strconv.Atoi(vs[1])
	v3, _ := strconv.Atoi(vs[2])
	return v1, v2, v3
}

type point struct {
	hi, wi int
}

var seen = map[point]int{}

func dfs2(m [][]byte, ecans map[point]int, p point, e int) {
	for _, d := range []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		np := point{p.hi + d.hi, p.wi + d.wi}
		// 探索済みかつその時のエネルギーが現在のエネルギーより小さい場合はスキップ
		if v := seen[np]; v >= 0 && v <= e-1 {
			continue
		}
		if np.hi < 0 || np.hi >= len(m) || np.wi < 0 || np.wi >= len(m[0]) {
			continue
		}
		if m[np.hi][np.wi] == '#' {
			continue
		}
		dfs(m, ecans, np, e-1)
	}
}

func dfs(m [][]byte, ecans map[point]int, p point, e int) {
	charged := ecans[p]
	if charged > e {
		e = charged
	}

	seen[p] = e

	if e > 0 {
		dfs2(m, ecans, p, e)
	}
}

func solve(r io.Reader) bool {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	h, w := scan2(sc)

	var sp, tp point
	m := make([][]byte, h)
	for hi := 0; hi < h; hi++ {
		sc.Scan()
		m[hi] = []byte(sc.Text())
		for wi := 0; wi < w; wi++ {
			if m[hi][wi] == 'S' {
				sp = point{hi, wi}
			}
			if m[hi][wi] == 'T' {
				tp = point{hi, wi}
			}
		}
	}

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	ecans := map[point]int{}
	for i := 0; i < n; i++ {
		hi, wi, e := scan3(sc)
		ecans[point{hi - 1, wi - 1}] = e
	}

	for hi := 0; hi < h; hi++ {
		for wi := 0; wi < w; wi++ {
			seen[point{hi, wi}] = -1
		}
	}

	// 深さ優先探索
	dfs(m, ecans, sp, 0)
	if v := seen[tp]; v == -1 {
		return false
	}
	return true
}
