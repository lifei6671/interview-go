package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	rmx      sync.RWMutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.rmx.Lock()
				for k, v := range o.visitIPs {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(o.visitIPs, k)
					}
				}
				o.rmx.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}
func (o *Ban) visit(ip string) bool {
	o.rmx.RLock()
	if _, ok := o.visitIPs[ip]; ok {
		o.rmx.RUnlock()
		return true
	}
	o.rmx.RUnlock()
	o.rmx.Lock()
	o.visitIPs[ip] = time.Now()
	o.rmx.Unlock()
	return false
}
func main() {
	success := 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ban := NewBan(ctx)

	wait := &sync.WaitGroup{}

	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}(j)
		}

	}
	wait.Wait()

	fmt.Println("success:", success)
}
