package main

import (
	"bufio"
	"io"
	"strconv"
)

// スタック(LIFO)による逆ポーランド記法の計算
// 数字ならスタックに積んで、演算子ならスタックから2つ取り出して計算して結果をスタックに積む

func solve2(r io.Reader) int {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	var stack []int
	for sc.Scan() {
		s := sc.Text()
		if s == "+" || s == "-" || s == "*" || s == "/" {
			if len(stack) < 2 {
				panic("stack underflow")
			}
			r := stack[len(stack)-1]
			l := stack[len(stack)-2]
			var ans int
			switch s {
			case "+":
				ans = l + r
			case "-":
				ans = l - r
			case "*":
				ans = l * r
			case "/":
				ans = l / r
			}
			stack[len(stack)-2] = ans
			stack = stack[:len(stack)-1]
			continue
		}
		n, _ := strconv.Atoi(s)
		stack = append(stack, n)
	}
	if len(stack) != 1 {
		panic("calculation failed")
	}
	return stack[0]
}
