package main

import (
	"strconv"
	"fmt"
	"strings"
)

func run(s string) (x, y int) {
	cmdList := resolveCmd(s)

	face := "Y"
	for _,c := range cmdList {
		if c == "L" {
			if face == "Y" {
				face = "-X"
			}else if face == "-Y"{
				face = "X"
			}else if face == "X" {
				face = "Y"
			}else {
				face = "-Y"
			}
		}else if c == "R" {
			if face == "Y" {
				face = "X"
			}else if face == "-Y"{
				face = "-X"
			}else if face == "X" {
				face = "-Y"
			}else {
				face = "Y"
			}
		}else if c == "F" {
			if face == "Y" {
				y += 1
			}else if face == "-Y" {
				y -= 1
			}else if face == "X" {
				x+=1
			}else {
				x-=1
			}
		}else if c == "B" {
			if face == "Y" {
				y-=1
			}else if face == "-Y" {
				y+=1
			}else if face == "X" {
				x+=1
			}else {
				x-=1
			}
		}
	}
	return
}

func resolveCmd(s string) ([]string){

	cmdList := make([]string,0)
	repeatCount := 0
	isStart := false
	tempCmd := ""

	for _,v := range s {
		ns := string(v)
		//如果是字符串，则标识下一步是重复步骤
		if ns >= "0" && ns <= "9" {
			t,_ := strconv.Atoi(ns);
			repeatCount = t
		}else if ns == "(" {
			isStart = true
		}else if ns == ")" {
			c := strings.Repeat(tempCmd,repeatCount)

			tempList := make([]string,strings.Count(c,""))

			for i,v1 := range c {
				tempList[i] = string(v1)
			}

			//当解析结束时，重复命令并保存到列表中
			cmdList = append(cmdList,tempList...)
			isStart = false
			repeatCount = 0
			tempCmd = ""
		}else if isStart{
			tempCmd += ns
		}else{
			cmdList = append(cmdList,ns)
		}
	}
	return cmdList
}

func main() {
	s := "R2(LF)"

	fmt.Println(run(s))
}
