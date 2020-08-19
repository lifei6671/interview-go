package main

import "fmt"

func main() {
	haystack := "abcabdabceabd"
	needle := "abceabb"
	next := getNext(needle)
	fmt.Println(next)
	index := kmpSearch(haystack, needle, next)
	fmt.Println(index)
}

func kmpSearch(haystack, needle string, next []int) int {
	l1 := len(haystack)
	l2 := len(needle)
	i, j := 0, 0
	for i < l1 && j < l2 {
		if j == -1 || haystack[i] == needle[j] {
			j++
			i++
		} else {
			j = next[j]
		}
	}
	if j == l2 {
		return i - j
	}
	return -1
}

func getNext(needle string) []int {
	next := make([]int, len(needle))
	next[0] = -1
	i, j := 0, -1
	for i < len(needle)-1 {
		if j == -1 || needle[i] == needle[j] {
			//真前缀
			i += 1
			//真后缀
			j += 1
			next[i] = j
		} else {
			//真后缀回溯
			j = next[j]
		}
	}
	return next
}
