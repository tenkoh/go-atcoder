package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type unionFind struct {
	// 各頂点の親の頂点番号。その頂点自身が根の場合は -1
	par []int
	// 各頂点の属するグループの頂点数
	siz []int
}

func newUnionFind(n int) *unionFind {
	uf := &unionFind{
		par: make([]int, n),
		siz: make([]int, n),
	}
	for i := range uf.par {
		uf.par[i] = -1
		uf.siz[i] = 1
	}
	return uf
}

// 頂点xの根を求める。同時に経路圧縮を行う。
func (uf *unionFind) root(x int) int {
	if uf.par[x] == -1 {
		return x
	}
	uf.par[x] = uf.root(uf.par[x])
	return uf.par[x]
}

// 頂点xを含むグループと頂点yを含むグループを併合する。
func (uf *unionFind) unite(x, y int) bool {
	x = uf.root(x)
	y = uf.root(y)
	if x == y {
		return false
	}
	// サイズが小さいグループを大きいグループに併合する
	if uf.siz[x] < uf.siz[y] {
		x, y = y, x
	}
	uf.par[y] = x
	uf.siz[x] += uf.siz[y]
	return true
}

// x が属するグループのサイズを返す。
func (uf *unionFind) size(x int) int {
	return uf.siz[uf.root(x)]
}

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type edge struct {
	a, b int
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n, m := scan1(sc), scan1(sc)

	edges := make([]edge, m)
	for i := range edges {
		edges[i] = edge{scan1(sc) - 1, scan1(sc) - 1}
	}

	// ここから辺を1つずつ取り除いたグラフを作り、各グラフが連結かどうかを調べる
	var ans int
	for i := 0; i < m; i++ {
		// 連結なグラフなので頂点数は辺を1つ取り除いても変わらない
		uf := newUnionFind(n)
		for j := 0; j < m; j++ {
			// 橋を取り除く操作に相当する
			if i == j {
				continue
			}
			uf.unite(edges[j].a, edges[j].b)
		}
		var s int
		for j := 0; j < n; j++ {
			if uf.root(j) == j {
				s++
			}
		}
		if s != 1 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(solve(os.Stdin))
}
