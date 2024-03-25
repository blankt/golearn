package main

// f(n) = f(n-1) + f(n-2)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	n1, n2 := 0, 1
	result := n2
	for i := 2; i <= n; i++ {
		result = n1 + n2
		n1 = n2
		n2 = result
	}
	return result
}
