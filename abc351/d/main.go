package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	hi, wi int
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, 10000), 1000000)

	var h, w int
	sc.Scan()
	hw := strings.Split(sc.Text(), " ")
	h, _ = strconv.Atoi(hw[0])
	w, _ = strconv.Atoi(hw[1])

	m := make([][]byte, h)
	for i := 0; i < h; i++ {
		sc.Scan()
		m[i] = []byte(sc.Text())
	}

	// 座標関係をグラフに直す
	// マスが磁石('#')である場合はスキップ
	// マス磁石ではないが、上下左右のどこかに磁石がある場合は辺を張らない
	// それ以外は隣接する磁石でないマスに辺を張る
	var g = map[coord][]coord{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if m[i][j] == '#' {
				continue
			}
			canMove := true
			candidates := make([]coord, 0, 4)
			for _, d := range []coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				ni, nj := i+d.hi, j+d.wi
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if m[ni][nj] == '#' {
					canMove = false
					break
				}
				candidates = append(candidates, coord{ni, nj})
			}
			if canMove {
				g[coord{i, j}] = candidates
			}
		}
	}

	maxDist := 1
	// 計算量削減のためにスキップ可能な頂点を記録
	canSkip := map[coord]struct{}{}
	// 各頂点から辿り着ける頂点の数を、その頂点自身も含めて数える
	for start := range g {
		if _, ok := canSkip[start]; ok {
			continue
		}
		visited := map[coord]bool{}
		var dfs func(coord)
		dfs = func(v coord) {
			visited[v] = true
			hasMagnets := false
			for _, d := range []coord{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				ni, nj := v.hi+d.hi, v.wi+d.wi
				if ni < 0 || ni >= h || nj < 0 || nj >= w {
					continue
				}
				if m[ni][nj] == '#' {
					hasMagnets = true
					break
				}
			}
			if !hasMagnets {
				canSkip[v] = struct{}{}
			}
			for _, u := range g[v] {
				if visited[u] {
					continue
				}
				dfs(u)
			}
		}
		dfs(start)
		if len(visited) > maxDist {
			maxDist = len(visited)
		}
	}
	return maxDist
}

func main() {
	fmt.Println(solve(os.Stdin))
}
