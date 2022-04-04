## 交替打印数字和字母

**问题描述**

使用两个 `goroutine` 交替打印序列，一个 `goroutine` 打印数字， 另外一个 `goroutine` 打印字母， 最终效果如下：

```bash
12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
```

**解题思路**

问题很简单，使用 channel 来控制打印的进度。使用两个 channel ，来分别控制数字和字母的打印序列， 数字打印完成后通过 channel 通知字母打印, 字母打印完成后通知数字打印，然后周而复始的工作。

**源码参考**

```go
	letter,number := make(chan bool),make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for{
			select {
			case <-letter:
				if i >= 'Z' {
					wait.Done()
					return
				}

				fmt.Print(string(i))
				i++
				fmt.Print(string(i))
				i++
				number <- true
			}

		}
	}(&wait)
	number<-true
	wait.Wait()
```

**源码解析**

这里用到了两个`channel`负责通知，letter负责通知打印字母的goroutine来打印字母，number用来通知打印数字的goroutine打印数字。wait用来等待字母打印完成后退出循环。


也可以分别使用三个 channel 来控制数字，字母以及终止信号的输入.

```go

package main

import "fmt"

func main() {
	number := make(chan bool)
	letter := make(chan bool)
	done := make(chan bool)

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	go func() {
		j := 'A'
		for {
			select {
			case <-letter:
				if j >= 'Z' {
					done <- true
				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					number <- true
				}
			}
		}
	}()

	number <- true

	for {
		select {
		case <-done:
			return
		}
	}
}
```
