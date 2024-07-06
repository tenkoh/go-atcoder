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

type Cube struct {
	xmin, ymin, zmin, xmax, ymax, zmax int
}

func chmin[T int | int64 | float32 | float64](a *T, b T) {
	if *a > b {
		*a = b
	}
}

func NewCube(sc *bufio.Scanner) *Cube {
	a := scan1(sc)
	b := scan1(sc)
	c := scan1(sc)
	d := scan1(sc)
	e := scan1(sc)
	f := scan1(sc)

	chmin(&a, d)
	chmin(&b, e)
	chmin(&c, f)
	return &Cube{a, b, c, d, e, f}
}

func (c *Cube) HasProduct(other *Cube) bool {
	if c.xmax <= other.xmin || c.xmin >= other.xmax {
		return false
	}
	if c.ymax <= other.ymin || c.ymin >= other.ymax {
		return false
	}
	if c.zmax <= other.zmin || c.zmin >= other.zmax {
		return false
	}
	return true
}

func solve(r io.Reader) bool {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	c1 := NewCube(sc)
	c2 := NewCube(sc)
	return c1.HasProduct(c2)
}

func main() {
	if solve(os.Stdin) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
