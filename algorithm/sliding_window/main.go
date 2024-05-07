package main

import "fmt"

/*
*
滑动窗口算法
思想：在给定窗口大小的数组或者字符串上执行要求的操作,可以将一部分问题中的嵌套循环转变为一个单循环，可以减少时间复杂度。
思路：
 1. 初始化窗口
 2. 扩大窗口直至找到满足条件的结果
 3. 缩小窗口寻找最优解
 4. 直至此次循环结束
*/
func main() {
	//s := "pwwkew"
	//fmt.Println(lengthOfLongestSubstring(s))
	//fmt.Println(lengthOfLongestSubstring1(s))

	s := "abaacbabc"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
