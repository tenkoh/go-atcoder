package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
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

func bfs(m [][]byte, ecans map[point]int, sp point) {
	// 最初にエネルギーが得られなかった場合は終了
	e := ecans[sp]
	seen[sp] = e
	if e == 0 {
		return
	}
	q := list.New()
	q.PushBack(sp)

	for q.Len() > 0 {
		p := q.Remove(q.Front()).(point)
		for _, d := range []point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			np := point{p.hi + d.hi, p.wi + d.wi}
			if np.hi < 0 || np.hi >= len(m) || np.wi < 0 || np.wi >= len(m[0]) {
				continue
			}
			if m[np.hi][np.wi] == '#' {
				continue
			}
			// 次の点でのエネルギー
			e := seen[p] - 1
			charged, exist := ecans[np]
			if exist && charged > e {
				e = charged
			}
			if e < 0 {
				continue
			}
			// 未探索またはエネルギーがより大きくできるときは探索
			if v := seen[np]; v < 0 || v < e {
				seen[np] = e
				q.PushBack(np)
				continue
			}
		}
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
			seen[point{hi, wi}] = -1
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

	// 深さ優先探索
	// dfs(m, ecans, sp, 0)
	bfs(m, ecans, sp)
	if v := seen[tp]; v == -1 {
		return false
	}
	return true
}

func main() {
	reachable := solve(os.Stdin)
	if reachable {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
