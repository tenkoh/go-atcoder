package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var ss = "wbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbwbwbwwbwbwwbw"

func solve(w, b int) bool {
	// 尺取法で解く
	// left, right を進めていってwの数、またはbの数が足りなければrightを進め、
	// ちょうど数が指定した数になればtrueを返して終わり
	// 途中でwの数またはbの数が指定した数を超えたらleftを進める
	// rightが最後まで進んだらfalseを返す
	var wcount, bcount int
	var right int
	for left := 0; left < len(ss); left++ {
		for right < len(ss) && (wcount < w || bcount < b) {
			if ss[right] == 'w' {
				wcount++
			} else {
				bcount++
			}
			right++
		}
		if wcount == w && bcount == b {
			return true
		}
		// 最後までrightを進めたのに条件を達成しなかったときはfalse
		if right == len(ss) {
			return false
		}
		if left == right {
			right++
		} else {
			if ss[left] == 'w' {
				wcount--
			} else {
				bcount--
			}
		}
	}
	return false
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	w, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	if solve(w, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
