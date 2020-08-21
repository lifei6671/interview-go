package main

import "fmt"

func main() {
	arr := []int{8, 12, 46, 12, 2, 4, 15, 3, 9, 44, 5, 6, 1, 59, 2}
	fmt.Println(insertionSort(arr))
}

func insertionSort(arr []int) []int {
	l := len(arr)
	if l == 0 {
		return arr
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
