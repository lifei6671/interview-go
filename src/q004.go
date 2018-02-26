package main

import (
	"strings"
	"fmt"
)

func isRegroup(s1,s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))

	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2{
		return false
	}

	for _,v := range s1 {
		if strings.Count(s1,string(v)) != strings.Count(s2,string(v)) {
			return false
		}
	}
	return true
}

func main() {
	s1 := "This is golang"
	s2 := "gnalog si sihT"
	fmt.Println(isRegroup(s1, s2))

	s3 := "Here you are"
	s4 := "Are you here"
	fmt.Println(isRegroup(s3, s4))

	s5 := "This is golang1.1"
	s6 := "This is golang1"
	fmt.Println(isRegroup(s5, s6))
}
