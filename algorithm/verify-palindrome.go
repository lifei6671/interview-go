package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
	s = "race a car"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	s = strings.ToLower(s)
	if len(s) == 2 {
		return s[0] == s[1]
	}
	left := 0
	right := len(s) - 1
	for left < right {
		if !((s[left] >= 'a' && s[left] <= 'z') || (s[left] >= '0' && s[left] <= '9')) {
			left++
			continue
		}
		if !((s[right] >= 'a' && s[right] <= 'z') || (s[right] >= '0' && s[right] <= '9')) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
