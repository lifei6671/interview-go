package main

import "fmt"

/**
题目描述
请实现有重复数字的有序数组的二分查找。
输出在数组中第一个大于等于查找值的位置，如果数组中不存在这样的数，则输出数组长度加一。

示例
5,4,[1,2,4,4,5]

输出：3
*/
func main() {
	arr := []int{3, 3, 4, 4, 4, 5, 6, 6, 6, 7, 8, 8, 12, 13, 15, 16, 21, 21, 22, 24, 24, 27, 28, 32, 34, 35, 35, 36, 36, 39, 40, 41, 41, 42, 44, 44, 45, 45, 47, 47, 47, 47, 48, 48, 50, 51, 51, 53, 53, 53, 54, 54, 54, 56, 56, 57, 59, 60, 60, 60, 60, 61, 62, 63, 65, 65, 65, 65, 67, 67, 68, 70, 71, 71, 74, 75, 75, 79, 81, 84, 84, 86, 86, 87, 90, 90, 90, 90, 91, 92, 93, 94, 94, 94, 95, 97, 97, 98, 98, 99}
	mid := upper_bound_(100, 97, arr)
	fmt.Println(mid)
}

/**
 * 二分查找
 * @param n int整型 数组长度
 * @param v int整型 查找值
 * @param a int整型一维数组 有序数组
 * @return int整型
 */
func upper_bound_(n int, v int, a []int) int {
	// write code here
	left := 0
	right := n - 1
	for left < right {
		mid := left + (right-left)/2
		if a[mid] >= v {
			if mid == 0 || a[mid-1] < v {
				return mid + 1
			}
			right = mid
		} else {
			left = mid + 1
		}
	}
	return n + 1
}
