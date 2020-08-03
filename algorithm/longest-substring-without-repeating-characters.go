package main

import "fmt"

//题目描述
//给定一个字符串，找出最长的不具有重复字符的子串的长度。例如，“abcabcbb”不具有重复字符的最长子串是“abc”，长度为3。对于“bbbbb”，最长的不具有重复字符的子串是“b”，长度为1。
func main() {
	s := "wlrbbmqbhcdarzowkkyhiddqscdxrjmowfrxsjybldbefsarcbynecdyggxxpklorellnmpapqfwkhopkmco"
	index := lengthOfLongestSubstring(s)
	fmt.Printf("正确输出 23 用例输出 %d", index)
}

func lengthOfLongestSubstring(s string) int {
	// write code here
	win := make(map[string]struct{})
	size := len(s)
	i := 0
	j := 0
	maxLen := 0
	for i < size && j < size {
		if _, ok := win[string(s[j])]; !ok {
			win[string(s[j])] = struct{}{}
			j++
			if m := j - i; m > maxLen {
				maxLen = m
			}
		} else {
			delete(win, string(s[i]))
			i++
		}
	}
	return maxLen
}
