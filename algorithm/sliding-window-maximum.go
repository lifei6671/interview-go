package main

import "fmt"

//给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//返回滑动窗口中的最大值所构成的数组。
//输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
//输出: [3,3,5,5,6,7]
//解释:
//
//  滑动窗口的位置                最大值
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
func main() {
	arr := []int{1, 3}
	fmt.Println(arr[0:1])
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	ret := maxSlidingWindow2(nums, k)
	fmt.Println(ret)
}

func maxSlidingWindow(nums []int, k int) []int {
	l1 := len(nums)
	ret := make([]int, 0)
	if l1 == 0 || k == 0 {
		return ret
	}
	index := 0
	for index < l1 {
		m := nums[index]
		if index > l1-k {
			break
		}
		for j := index + 1; j < index+k; j++ {
			if m < nums[j] {
				m = nums[j]
			}
		}
		ret = append(ret, m)
		index++
	}
	return ret
}
func maxSlidingWindow2(nums []int, k int) []int {
	ret := make([]int, 0)
	if len(nums) == 0 {
		return ret
	}
	var queue []int
	for i := range nums {
		for i > 0 && (len(queue) > 0) && nums[i] > queue[len(queue)-1] {
			//将比当前元素小的元素祭天
			queue = queue[:len(queue)-1]
		}
		//将当前元素放入queue中
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			//维护队列，保证其头元素为当前窗口最大值
			queue = queue[1:]
		}
		if i >= k-1 {
			//放入结果数组
			ret = append(ret, queue[0])
		}
	}
	return ret
}
