package main

import "sort"

func solve(as []int) []int {
	sorted := append([]int{}, as...)
	sort.Ints(sorted)
	// 値:インデックスのマップ
	m := map[int]int{}
	for i, v := range sorted {
		m[v] = i
	}
	ans := make([]int, len(as))
	// 値をインデックスに変換
	for i, v := range as {
		ans[i] = m[v]
	}
	return ans
}
