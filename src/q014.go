package main

import (
	"context"
	"fmt"
)

type query func(string) string

func exec(name string, vs ...query) string {
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	fn := func(i int) {
		select {
		case ch <- vs[i](name):
		case <-ctx.Done():
			return
		}
	}
	for i := range vs {
		go fn(i)
	}
	defer cancel()
	return <-ch
}

func main() {
	ret := exec("111", func(n string) string {
		return n + "func1"
	}, func(n string) string {
		return n + "func2"
	}, func(n string) string {
		return n + "func3"
	}, func(n string) string {
		return n + "func4"
	})
	fmt.Println(ret)
}
