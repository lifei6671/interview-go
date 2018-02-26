package main

import (
	"fmt"
)

func reverString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	if l > 5000 {
		return string(str), false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-1-i] = str[l-1-i], str[i]
	}
	return string(str), true
}

func main() {
	s1 := "This is golang"
	fmt.Println(reverString(s1))

	s2 := "gnalog si sihT"
	fmt.Println(reverString(s2))
}