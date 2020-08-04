package main

import "fmt"

//两个数组的交集
//给定两个数组，编写一个函数来计算它们的交集。
//示例 1:
//
//输入: nums1 = [1,2,2,1], nums2 = [2,2]
//
//输出: [2,2]
//示例 2:
//
//输入: nums1 = [4,9,5,9], nums2 = [9,4,9,8,4]
//
//输出: [4,9,9]
//说明：
//
//输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//我们可以不考虑输出结果的顺序。
//进阶:
//
//如果给定的数组已经排好序呢？将如何优化你的算法呢？
//思路：设定两个为0的指针，比较两个指针的元素是否相等。如果指针的元素相等，我们将两个指针一起向后移动，并且将相等的元素放入空白数组。
func main() {
	nums1 := []int{4, 9, 5, 9}
	nums2 := []int{9, 4, 8, 4, 9, 5, 5}

	fmt.Println(intersect(nums1, nums2))
}

//无序数组
func intersect(nums1 []int, nums2 []int) []int {
	m0 := make(map[int]int)
	for _, i := range nums1 {
		m0[i] += 1
	}
	k := 0
	for _, v := range nums2 {
		if m0[v] > 0 {
			m0[v] -= 1
			//这里是复用切片
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}

//有序数组
func intersectSort(nums1 []int, nums2 []int) []int {
	return nil
}
