# 滑动窗口最大值

## 01、题目分析

### 第239题：滑动窗口最大值

给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。返回滑动窗口中的最大值。


返回滑动窗口中的最大值所构成的数组。


示例:

>输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
>输出: [3,3,5,5,6,7] 
>解释: 
>
>  滑动窗口的位置                最大值
>---------------               -----
>[1  3  -1] -3  5  3  6  7       3
> 1 [3  -1  -3] 5  3  6  7       3
> 1  3 [-1  -3  5] 3  6  7       5
> 1  3  -1 [-3  5  3] 6  7       5
> 1  3  -1  -3 [5  3  6] 7       6
> 1  3  -1  -3  5 [3  6  7]      7

## 02、题目分析

本题对于题目没有太多需要额外说明的，应该都能理解，直接进行分析。我们很容易想到，可以通过遍历所有的滑动窗口，找到每一个窗口的最大值，来进行暴力求解。那一共有多少个滑动窗口呢，小学题目，可以得到共有 L-k+1 个窗口。


假设 `nums = [1,3,-1,-3,5,3,6,7]`，和 `k = 3`，窗口数为`6`:

![](data:image/jpeg;base64,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)

根据分析，直接完成代码：

```go
func maxSlidingWindow(nums []int, k int) []int {
	l1 := len(nums)
	index := 0
	ret := make([]int, 0)
	for index < l1 {
		m := nums[index]
		if index > l1 - k {
			break
		}
		for j := index + 1; j < index + k; j++ {
			if m < nums[j] {
				m = nums[j]
			}
		}
		ret = append(ret,m)
		index++
	}
	return ret
}
```

## 03、线性题解

这里不卖关子，其实这道题比较经典，我们可以采用队列，DP，堆等方式进行求解，所有思路的主要源头应该都是在窗口滑动的过程中，如何更快的完成查找最大值的过程。但是最典型的解法还是使用双端队列。具体怎么来求解，一起看一下。


首先，我们了解一下，什么是双端队列：是一种具有队列和栈的性质的数据结构。双端队列中的元素可以从两端弹出或者插入。

![](../../images/3.395c1e89.jpg)

我们可以利用双端队列来实现一个窗口，目的是让该窗口可以做到张弛有度（汉语博大精深，也就是长度动态变化。其实用游标或者其他解法的目的都是一样的，就是去维护一个可变长的窗口）

然后我们再做一件事，只要遍历该数组，同时**在双端队列的头去维护当前窗口的最大值（在遍历过程中，发现当前元素比队列中的元素大，就将原来队列中的元素祭天），在整个遍历的过程中我们再记录下每一个窗口的最大值到结果数组中。**最终结果数组就是我们想要的，整体图解如下。

假设 `nums = [1,3,-1,-3,5,3,6,7]`，和 `k = 3`:

![](../../images/4.f3a9aa62.jpg)

根据分析，得出代码：

```go
func maxSlidingWindow2(nums []int, k int) []int {
	ret := make([]int,0)
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
```