package main

import "fmt"

//字符串中的第一个唯一字符
//给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1 。
//案例:
//
//s = "leetcode"
//返回 0.
//
//s = "loveleetcode",
//返回 2.
//注意事项： 您可以假定该字符串只包含小写字母。
func main() {
	s := "loveleetcode"
	i := firstUniqueChar(s)
	fmt.Println(i)
}

func firstUniqueChar(s string) int {
	var arr [26]int
	//第一次遍历计算所有字符出现的最后位置
	for i, k := range s {
		arr[k-'a'] = i
	}
	for i, k := range s {
		if arr[k-'a'] == i {
			return i
		}
	}
	return -1
}
