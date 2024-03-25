package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

/**
递归或者动态规划
递归要超时
动态规划状态转移方程:f(n) = f(n-1) + f(n-2)
*/

func climbStairs(n int) int {
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r += p
	}
	return r
}
