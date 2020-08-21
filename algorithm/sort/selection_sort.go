package main

import "fmt"

func main() {
	arr := []int{8, 12, 46, 12, 2, 4, 15, 3, 9, 44, 5, 6, 1, 59, 2}
	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return arr
	}
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
