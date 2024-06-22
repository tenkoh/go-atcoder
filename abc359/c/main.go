package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(sx, sy, tx, ty int) int {
	// 初期化。
	if sx < tx {
		if sy != ty && (sx+sy)%2 == 0 {
			sx++
		}
		if sy == ty && (sx+sy)%2 != 0 {
			sx--
		}
		if (tx+ty)%2 != 0 {
			tx--
		}
	} else if sx > tx {
		if sy != ty && (sx+sy)%2 != 0 {
			sx--
		}
		if sy == ty && (sx+sy)%2 == 0 {
			sx++
		}
		if (tx+ty)%2 == 0 {
			tx++
		}
	}

	// ty-txの数だけカウントを増やす
	var count int
	dy := ty - sy
	if dy < 0 {
		dy *= -1
	}
	count += dy

	if sx == tx {
		return count
	}

	dx := tx - sx
	if dx < 0 {
		dx *= -1
	}
	if dx < count {
		return count
	}
	return count + dx/2
}

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1000000)
	sx := scan1(sc)
	sy := scan1(sc)
	tx := scan1(sc)
	ty := scan1(sc)
	fmt.Println(solve(sx, sy, tx, ty))
}
