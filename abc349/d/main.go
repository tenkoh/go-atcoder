package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type s struct {
	l, r int
}

// nを超えない最大の2のべき乗を返す
func pow2(n int) int {
	res := 1
	for res*2 <= n {
		res *= 2
	}
	return res
}

func solve(whole s) []s {
	l := whole.l
	r := l + 1
	ans := []s{}
	for r <= whole.r {
		// 左端が奇数の場合はj=l, i=0。r=j+1
		if l%2 == 1 {
			ans = append(ans, s{l, r})
			l = r
			r = l + 1
			continue
		}

		// 左端が0の場合はj=0. rは行けるだけ行く
		if l == 0 {
			pow := pow2(whole.r)
			ans = append(ans, s{0, pow})
			l = pow
			r = l + 1
			continue
		}

		pow := pow2(l)
		if whole.r-l < pow {
			pow = pow2(whole.r - l)
		}
		j := l / pow
		r = pow * (j + 1)
		ans = append(ans, s{l, r})
		l = r
		r = l + 1
	}
	return ans
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), 0)
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	areas := solve(s{l: l, r: r})
	fmt.Println(len(areas))
	bw := bufio.NewWriter(os.Stdout)
	for _, v := range areas {
		fmt.Fprintln(bw, v.l, v.r)
	}
	bw.Flush()
}
