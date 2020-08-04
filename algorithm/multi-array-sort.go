package main

import (
	"fmt"
	"sort"
)

//按照二维数组的第一列升序排序，第二列降序排列
func main() {
	nums := [][]int{{1, 2, 3}, {2, 4, 4}, {2, 3, 1}, {1, 3, 1}}
	column := 2
	fmt.Printf("根据第%d列降序排序：\n", column)
	fmt.Println("原数组 ->", nums)
	result := sortArray(nums, column, true)
	fmt.Println("排序后 ->", result)
}

func sortArray(arr [][]int, column int, desc bool) [][]int {
	intArr := &IntArray{
		arr:    arr,
		column: column,
		desc:   desc,
	}
	sort.Sort(intArr)
	return intArr.arr
}

type IntArray struct {
	arr    [][]int
	column int
	desc   bool
}

func (arr *IntArray) Len() int {
	return len(arr.arr)
}
func (arr *IntArray) Less(i, j int) bool {
	if arr.arr[i][arr.column] < arr.arr[j][arr.column] {
		return !arr.desc
	} else if arr.arr[i][arr.column] > arr.arr[j][arr.column] {
		return arr.desc
	}
	return true
}
func (arr *IntArray) Swap(i, j int) {
	arr.arr[i], arr.arr[j] = arr.arr[j], arr.arr[i]
}
