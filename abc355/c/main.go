package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func solve(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	n := scan1(sc)
	t := scan1(sc)

	// ビンゴ判定用のmapを作る.オープンされる数字はすべて異なるため、
	// 縦・横・斜めの各列内のヒット数をカウントすればビンゴ判定可能
	yoko := make(map[int]int)
	tate := make(map[int]int)
	naname1 := 0 // 左上から右下への斜め
	judgeNaname1 := make(map[int]struct{})
	naname2 := 0 // 右上から左下への斜め
	judgeNaname2 := make(map[int]struct{})

	// 斜め判定用のmapを作る
	for i := 1; i <= n; i++ {
		judgeNaname1[n*(i-1)+i] = struct{}{}
		judgeNaname2[n*(i-1)+(n-i+1)] = struct{}{}
	}

	for i := 1; i <= t; i++ {
		a := scan1(sc)

		yoko[(a-1)/n]++
		if yoko[(a-1)/n] == n {
			return i
		}

		tate[a%n]++
		if tate[a%n] == n {
			return i
		}

		// 斜めのカウント
		if _, ok := judgeNaname1[a]; ok {
			naname1++
		}
		if naname1 == n {
			return i
		}

		if _, ok := judgeNaname2[a]; ok {
			naname2++
		}
		if naname2 == n {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(solve(os.Stdin))
}
