package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 5)
	data := []int{1, 2, 3, 10, 999, 8, 345, 7, 98, 33, 66, 77, 88, 68, 96}
	dataLen := len(data)
	size := 3
	target := 345
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan bool)
	wg := &sync.WaitGroup{}
	finishChan := make(chan struct{})
	for i := 0; i < dataLen; i += size {
		end := i + size
		if end >= dataLen {
			end = dataLen - 1
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			SearchTarget(ctx, data[i:end], target, resultChan)
		}()
	}
	go func() {
		wg.Wait()
		finishChan <- struct{}{}
	}()
	select {
	case <-timer.C:
		fmt.Fprintln(os.Stderr, "Timeout! Not Found")
		cancel()
	case <-resultChan:
		fmt.Fprintf(os.Stdout, "Found it!\n")
		cancel()
	case <- finishChan:
		fmt.Printf("[%d] not found in slice\n", target)
	}

	time.Sleep(time.Second * 2)
}

func SearchTarget(ctx context.Context, data []int, target int, resultChan chan bool) {
	for _, v := range data {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stdout, "Task cancelded! \n")
			return
		default:
		}
		// 模拟一个耗时查找，这里只是比对值，真实开发中可以是其他操作
		fmt.Fprintf(os.Stdout, "v: %d \n", v)
		time.Sleep(time.Millisecond * 1500)
		if target == v {
			resultChan <- true
			return
		}
	}

}
