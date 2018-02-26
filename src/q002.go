//判断字符串中字符是否全都不同
package main

import (
	"strings"
	"fmt"
)

func isUniqueString(s string) bool {
	if strings.Count(s,"") > 3000{
		return  false
	}
	for _,v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s,string(v)) > 1 {
			return false
		}
	}
	return true
}

func isUniqueString2(s string) bool {
	if strings.Count(s,"") > 3000{
		return  false
	}
	for _,v := range s {
		if v > 127 {
			return false
		}
		if strings.Index(s,string(v)) != strings.LastIndex(s,string(v)) {
			return false
		}
	}
	return true
}
func main() {
	str := "abfgasdgfdgfdfdfds"

	fmt.Println(isUniqueString2(str))

	str = "abcdefG"
	fmt.Println(isUniqueString2(str))
}