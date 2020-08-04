package main

import "fmt"

//现在有一个包含n个物体的数组，其中物体颜色为颜色为红色、白色或蓝色，请对这个数组进行排序，让相同颜色的物体相邻，颜色的顺序为红色，白色，蓝色。
//我们用0,1,2分别代表颜色红，白，蓝
//注意：
//本题要求你不能使用排序库函数
//扩展：
//一个非常直接的解法是两步的计数排序的算法
//首先：遍历一遍数组，记录0,1,2的数量，然后重写这个数组，先将0写入，再将1写入，再将2写入
//你能给出一个只用一步，并且能在常数级空间复杂度解决这个问题的算法吗？
func main() {
	nums := []int{0, 1, 2, 0, 1, 2}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(arr []int) {
	l := len(arr)
	if l == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
	} else if l > 2 {
		m := 0
		j := l - 1
		i := 0
		for i < l {
			if m > 2 {
				break
			}
			if arr[j] == m {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j++
			}
			if i == j {
				m = m + 1
				j = l - 1
				continue
			}
			j--
		}
	}
}
