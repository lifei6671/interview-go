package main

import (
	"fmt"
)

// 编写一个函数查找字符串数组的公共最长前缀，如果不存在则返回空
// 输入 ["flower","flow","flight"]
// 输出 fl
// 输入 【"dog","racecar","car"]
// 输出 ""
func main() {
	arr := []string{"fllower", "fllow", "fllight"}
	s := getPrefix(arr)
	fmt.Println(s)
	arr = []string{"dog", "racecar", "car"}
	s = getPrefix(arr)
	fmt.Println(s)
}

func getPrefix(arr []string) string {
	if len(arr) <= 1 {
		return ""
	}
	firstStr := arr[0]
	l := len(arr)
	for i := range firstStr {
		for j := 1; j < l; j++ {
			if arr[j][i] != firstStr[i] {
				return firstStr[:i]
			}
		}
	}
	return ""
}
