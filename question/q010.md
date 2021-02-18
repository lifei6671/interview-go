# 实现阻塞读且并发安全的map

GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接口：

```go
type sp interface {
    Out(key string, val interface{})  //存入key /val，如果该key读取的goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
    Rd(key string, timeout time.Duration) interface{}  //读取一个key，如果key不存在阻塞，等待key存在或者超时
}
```

**解析：**

看到阻塞协程第一个想到的就是`channel`，题目中要求并发安全，那么必须用锁，还要实现多个`goroutine`读的时候如果值不存在则阻塞，直到写入值，那么每个键值需要有一个阻塞`goroutine` 的 `channel`。

[实现如下：](../src/q010.go) 

```go
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
			value: val,
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
```
