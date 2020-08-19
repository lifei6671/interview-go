package main

import "fmt"

//题目描述
//给定一个字符串，找出最长的不具有重复字符的子串的长度。例如，“abcabcbb”不具有重复字符的最长子串是“abc”，长度为3。对于“bbbbb”，最长的不具有重复字符的子串是“b”，长度为1。
func main() {
	s := "abcabcbb"
	index := lengthOfLongestSubstring2(s)
	fmt.Printf("输出 %d", index)
}

//双指针法
func lengthOfLongestSubstring(s string) int {
	// write code here
	win := make(map[uint8]struct{})
	size := len(s)
	i := 0
	j := 0
	maxLen := 0
	for i < size && j < size {
		if _, ok := win[s[j]]; !ok {
			win[s[j]] = struct{}{}
			j++
			if m := j - i; m > maxLen {
				maxLen = m
			}
		} else {
			delete(win, s[i])
			i++
		}
	}
	return maxLen
}

func lengthOfLongestSubstring2(s string) int {
	win := make([]int, 256)
	size := len(s)
	left, right := 0, 0
	maxLen := 0
	for ; right < size; right++ {
		if win[s[right]] > left {
			left = win[s[right]]
		}
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
		win[s[right]] = right + 1
	}
	return maxLen
}
