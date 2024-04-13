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

// see: https://atcoder.jp/contests/abc349/editorial/9788
func f(l, r, L, R int) []s {
	if L <= l && r <= R {
		return []s{{l, r}}
	}
	m := (l + r) / 2
	if R <= m {
		return f(l, m, L, R)
	}
	if m <= L {
		return f(m, r, L, R)
	}
	return append(f(l, m, L, R), f(m, r, L, R)...)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 1024), 0)
	sc.Scan()
	l, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	areas := f(0, 1<<60, l, r)
	fmt.Println(len(areas))
	bw := bufio.NewWriter(os.Stdout)
	for _, v := range areas {
		fmt.Fprintln(bw, v.l, v.r)
	}
	bw.Flush()
}
