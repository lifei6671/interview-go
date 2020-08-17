package main

import "fmt"

//实现 strStr()
//实现 strStr() 函数。给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。
//示例 1:
//
//输入: haystack = "hello", needle = "ll"
//输出: 2
//示例 2:
//
//输入: haystack = "aaaaa", needle = "bba"
//输出: -1
// https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/303.html#_02%E3%80%81sunday-%E5%8C%B9%E9%85%8D
func main() {
	haystack := "Here is a little Hao"
	needle := "little"
	index := strStrSunday(haystack, needle)
	fmt.Println(index)
}

//暴力解法
func strStr(haystack, needle string) int {
	index := -1
	needleIndex := 0
	for i := range haystack {
		if haystack[i] == needle[needleIndex] {
			needleIndex++
			if needleIndex >= len(needle) {
				return index
			}
			if index == -1 {
				index = i
			}
			continue
		}
		index = -1
		needleIndex = 0
	}
	return index
}

//使用Sunday算法
func strStrSunday(haystack, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if haystack == needle {
		return 0
	}
	index := -1
	i := 0
	needleIndex := 0
	for i < len(haystack) {
		if haystack[i] == needle[needleIndex] {
			if index == -1 {
				index = i
			}
			i++
			needleIndex++
			if needleIndex >= len(needle) {
				break
			}
			continue
		}
		index = -1
		i = i + len(needle) - needleIndex
		if i >= len(haystack) {
			return index
		}
		offset := 0
		for j := len(needle) - 1; j > 0; j-- {
			if haystack[i] == needle[j] {
				offset = j
				break
			}
		}

		i = i - offset
		needleIndex = 0
	}
	return index
}
