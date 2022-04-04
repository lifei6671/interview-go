# 在 golang 协程和channel配合使用

> 写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。

**解析**

这是一道很简单的golang基础题目，实现方法也有很多种，一般想答让面试官满意的答案还是有几点注意的地方。

1. `goroutine` 在golang中式非阻塞的
2. `channel` 无缓冲情况下，读写都是阻塞的，且可以用`for`循环来读取数据，当管道关闭后，`for` 退出。
3.  golang 中有专用的`select case` 语法从管道读取数据。

示例代码如下：

```go
func main() {
    out := make(chan int)
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 0; i < 5; i++ {
            out <- rand.Intn(5)
        }
        close(out)
    }()
    go func() {
        defer wg.Done()
        for i := range out {
            fmt.Println(i)
        }
    }()
    wg.Wait()
}
```

如果不想使用 `sync.WaitGroup`, 也可以用一个 `done` channel.
```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			num, ok := <-random
			if ok {
				fmt.Println(num)
			} else {
				done <- true
			}
		}
	}()

	go func() {
		defer close(random)

		for i := 0; i < 5; i++ {
			random <- rand.Intn(5)
		}
	}()

	<-done
	close(done)
}
```
