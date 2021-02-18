/*
@Time : 2021/1/29 上午9:31
@Author : sunyujia
@File : q010
@Software: GoLand
*/
package main

import (
	"log"
	"sync"
	"time"
)

type sp interface {
	Out(key string, val interface{})                  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	item, ok := m.c[key]
	if !ok {
		m.c[key] = &entry{
			value:   val,
			isExist: true,
		}
		return
	}
	item.value = val
	if !item.isExist {
		if item.ch != nil {
			close(item.ch)
			item.ch = nil
		}
	}
	return
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.RLock()
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.RUnlock()
		return e.value
	} else if !ok {
		m.rmx.RUnlock()
		m.rmx.Lock()
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = e
		m.rmx.Unlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	} else {
		m.rmx.RUnlock()
		log.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			log.Println("协程超时 -> ", key)
			return nil
		}
	}
}

func main() {
	mapval := Map{
		c:   make(map[string]*entry),
		rmx: &sync.RWMutex{},
	}

	for i := 0; i < 10; i++ {
		go func() {
			val := mapval.Rd("key", time.Second*6)
			log.Println("读取值为->", val)
		}()
	}

	time.Sleep(time.Second * 3)
	for i := 0; i < 10; i++ {
		go func(val int) {
			mapval.Out("key", val)
		}(i)
	}

	time.Sleep(time.Second * 30)
}
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
