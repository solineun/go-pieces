package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	input1, _ := in.ReadString('\n')
	n, _ := strconv.Atoi(input1[:1])
	m, _ := strconv.Atoi(input1[2:])
	input2, _ := in.ReadString('\n')
	pts := []int{}
	for _, f := range strings.Fields(input2) {
		n, _ := strconv.Atoi(f)
		pts = append(pts, n)
	}
	fmt.Println(Points(n, m, pts))
}

func Points(n, m int, fpts []int) int {
	nn := make([]int, 0)
	ans := 0
	for _, fp := range fpts {
		if fp != 0 {
			ans += fp * fp
			nn = append(nn, fp)
		}
	}
	secpts := make([]int, 0)
	for i, n := range nn {
		l := i + 1
		r := l + n
		if r > len(nn) {
			r = len(nn)
		}
		secpts = append(secpts, nn[l:r]...)
	}
	for _, pt := range secpts {
		ans += pt
	}
	return ans
}