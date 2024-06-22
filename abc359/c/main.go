package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(sx, sy, tx, ty int) int {
	// 初期化：タイルの左端に寄せておく
	if (sx+sy)%2 != 0 {
		sx--
	}
	if (tx+ty)%2 != 0 {
		tx--
	}

	var count int
	count += absint(sy - ty)

	if sx == tx {
		return count
	}

	dx := absint(sx - tx)
	if dx <= count {
		return count
	}
	return count + (dx-count)/2
}

func absint[T int | int32 | int64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func scan1(sc *bufio.Scanner) int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sx := scan1(sc)
	sy := scan1(sc)
	tx := scan1(sc)
	ty := scan1(sc)
	fmt.Println(solve(sx, sy, tx, ty))
}
