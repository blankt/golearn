package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{2}
	dp := make(map[int]int)
	fmt.Println(coinChange2(coins, 1, &dp))
}

func coinChange2(coins []int, amount int, dp *map[int]int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	result := math.MaxInt

	for _, v := range coins {
		var res int
		result1, ok := (*dp)[amount-v]
		if ok {
			res = result1
		} else {
			res = coinChange2(coins, amount-v, dp)
			if res < 0 {
				continue
			}
		}

		result = min(result, res+1)
		(*dp)[amount] = result
	}

	if result == math.MaxInt {
		return -1
	} else {
		return result
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
