package main

import "fmt"

//输入 1111hhhh333hnn444nn
//输出
//_11111
//hhhh_333
//hnn_444
//nn
func main() {
	s := "1111hhhh333hnn444nn"
	covert(s)
}

func covert(s string) {
	if s == "" {
		return
	}
	l := len(s)
	for i, c := range s {
		if c >= '0' && c <= '9' {
			if i == 0 {
				fmt.Print("_")
			}
			fmt.Print(string(c))
			if i != l-1 && s[i+1] >= 'a' && s[i+1] <= 'z' {
				fmt.Println("")
			}
		}
		if c >= 'a' && c <= 'z' {
			fmt.Print(string(c))
			if i != l-1 && s[i+1] >= '0' && s[i+1] <= '9' {
				fmt.Print("_")
			}
		}
	}
}
