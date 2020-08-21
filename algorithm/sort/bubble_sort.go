package main

import "fmt"

func main() {
	arr := []int{8, 12, 46, 12, 2, 4, 15, 3, 9, 44, 5, 6, 1, 59, 2}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}
