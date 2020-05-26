# Goroutine调度策略

> 原文： [第三章 Goroutine调度策略（16）](https://mp.weixin.qq.com/s?__biz=MzU1OTg5NDkzOA==&mid=2247483801&idx=1&sn=ef7f872afccf148661cbd5a3d3b5b0a2&scene=19#wechat_redirect)


在调度器概述一节我们提到过，所谓的goroutine调度，是指程序代码按照一定的算法在适当的时候挑选出合适的goroutine并放到CPU上去运行的过程。这句话揭示了调度系统需要解决的三大核心问题：

- 调度时机：什么时候会发生调度？
- 调度策略：使用什么策略来挑选下一个进入运行的goroutine？
- 切换机制：如何把挑选出来的goroutine放到CPU上运行？

对这三大问题的解决构成了调度器的所有工作，因而我们对调度器的分析也必将围绕着它们所展开。

第二章我们已经详细的分析了调度器的初始化以及goroutine的切换机制，本章将重点讨论调度器如何挑选下一个goroutine出来运行的策略问题，而剩下的与调度时机相关的内容我们将在第4～6章进行全面的分析。

## 再探schedule函数

在讨论main goroutine的调度时我们已经见过schedule函数，因为当时我们的主要关注点在于main goroutine是如何被调度到CPU上运行的，所以并未对schedule函数如何挑选下一个goroutine出来运行做深入的分析，现在是重新回到schedule函数详细分析其调度策略的时候了。

[runtime/proc.go : 2467](https://github.com/golang/go/blob/bcda68447b31b86bc3829fca80454ca1a2a572e0/src/runtime/proc.go#L2551)

```go
// One round of scheduler: find a runnable goroutine and execute it.
// Never returns.
func schedule() {
    _g_ := getg()   //_g_ = m.g0

    ......

    var gp *g

    ......
   
    if gp == nil {
    // Check the global runnable queue once in a while to ensure fairness.
    // Otherwise two goroutines can completely occupy the local runqueue
    // by constantly respawning each other.
       //为了保证调度的公平性，每个工作线程每进行61次调度就需要优先从全局运行队列中获取goroutine出来运行，
       //因为如果只调度本地运行队列中的goroutine，则全局运行队列中的goroutine有可能得不到运行
        if _g_.m.p.ptr().schedtick%61 == 0 && sched.runqsize > 0 {
            lock(&sched.lock) //所有工作线程都能访问全局运行队列，所以需要加锁
            gp = globrunqget(_g_.m.p.ptr(), 1) //从全局运行队列中获取1个goroutine
            unlock(&sched.lock)
        }
    }
    if gp == nil {
        //从与m关联的p的本地运行队列中获取goroutine
        gp, inheritTime = runqget(_g_.m.p.ptr())
        if gp != nil && _g_.m.spinning {
            throw("schedule: spinning with local work")
        }
    }
    if gp == nil {
        //如果从本地运行队列和全局运行队列都没有找到需要运行的goroutine，
        //则调用findrunnable函数从其它工作线程的运行队列中偷取，如果偷取不到，则当前工作线程进入睡眠，
        //直到获取到需要运行的goroutine之后findrunnable函数才会返回。
        gp, inheritTime = findrunnable() // blocks until work is available
    }

    ......

    //当前运行的是runtime的代码，函数调用栈使用的是g0的栈空间
    //调用execte切换到gp的代码和栈空间去运行
    execute(gp, inheritTime)  
}
```

schedule函数分三步分别从各运行队列中寻找可运行的goroutine：

- 第一步，从全局运行队列中寻找goroutine。为了保证调度的公平性，每个工作线程每经过61次调度就需要优先尝试从全局运行队列中找出一个goroutine来运行，这样才能保证位于全局运行队列中的goroutine得到调度的机会。全局运行队列是所有工作线程都可以访问的，所以在访问它之前需要加锁。

- 第二步，从工作线程本地运行队列中寻找goroutine。如果不需要或不能从全局运行队列中获取到goroutine则从本地运行队列中获取。

- 第三步，从其它工作线程的运行队列中偷取goroutine。如果上一步也没有找到需要运行的goroutine，则调用findrunnable从其他工作线程的运行队列中偷取goroutine，findrunnable函数在偷取之前会再次尝试从全局运行队列和当前线程的本地运行队列中查找需要运行的goroutine。

下面我们先来看如何从全局运行队列中获取goroutine。

## 从全局运行队列中获取goroutine

从全局运行队列中获取可运行的goroutine是通过globrunqget函数来完成的，该函数的第一个参数是与当前工作线程绑定的p，第二个参数max表示最多可以从全局队列中拿多少个g到当前工作线程的本地运行队列中来。

[runtime/proc.go : 4663](https://github.com/golang/go/blob/bcda68447b31b86bc3829fca80454ca1a2a572e0/src/runtime/proc.go#L4996)

```go
// Try get a batch of G's from the global runnable queue.
// Sched must be locked.
func globrunqget(_p_ *p, max int32) *g {
    if sched.runqsize == 0 {  //全局运行队列为空
        return nil
    }

    //根据p的数量平分全局运行队列中的goroutines
    n := sched.runqsize / gomaxprocs + 1
    if n > sched.runqsize { //上面计算n的方法可能导致n大于全局运行队列中的goroutine数量
        n = sched.runqsize
    }
    if max > 0 && n > max {
        n = max   //最多取max个goroutine
    }
    if n > int32(len(_p_.runq)) / 2 {
        n = int32(len(_p_.runq)) / 2  //最多只能取本地队列容量的一半
    }

    sched.runqsize -= n

    //直接通过函数返回gp，其它的goroutines通过runqput放入本地运行队列
    gp := sched.runq.pop()  //pop从全局运行队列的队列头取
    n--
    for ; n > 0; n-- {
        gp1 := sched.runq.pop()  //从全局运行队列中取出一个goroutine
        runqput(_p_, gp1, false)  //放入本地运行队列
    }
    return gp
}
```

globrunqget函数首先会根据全局运行队列中goroutine的数量，函数参数max以及_p_的本地队列的容量计算出到底应该拿多少个goroutine，然后把第一个g结构体对象通过返回值的方式返回给调用函数，其它的则通过runqput函数放入当前工作线程的本地运行队列。这段代码值得一提的是，计算应该从全局运行队列中拿走多少个goroutine时根据p的数量（gomaxprocs）做了负载均衡。

如果没有从全局运行队列中获取到goroutine，那么接下来就在工作线程的本地运行队列中寻找需要运行的goroutine。

## 从工作线程本地运行队列中获取goroutine

从代码上来看，工作线程的本地运行队列其实分为两个部分，一部分是由p的runq、runqhead和runqtail这三个成员组成的一个无锁循环队列，该队列最多可包含256个goroutine；另一部分是p的runnext成员，它是一个指向g结构体对象的指针，它最多只包含一个goroutine。

从本地运行队列中寻找goroutine是通过`runqget`函数完成的，寻找时，代码首先查看`runnext`成员是否为空，如果不为空则返回runnext所指的goroutine，并把runnext成员清零，如果runnext为空，则继续从循环队列中查找goroutine。

[runtime/proc.go : 4825](https://github.com/golang/go/blob/bcda68447b31b86bc3829fca80454ca1a2a572e0/src/runtime/proc.go#L5192:1)

```go
// Get g from local runnable queue.
// If inheritTime is true, gp should inherit the remaining time in the
// current time slice. Otherwise, it should start a new time slice.
// Executed only by the owner P.
func runqget(_p_ *p) (gp *g, inheritTime bool) {
    // If there's a runnext, it's the next G to run.
    //从runnext成员中获取goroutine
    for {
        //查看runnext成员是否为空，不为空则返回该goroutine
        next := _p_.runnext  
        if next == 0 {
            break
        }
        if _p_.runnext.cas(next, 0) {
            return next.ptr(), true
        }
    }

    //从循环队列中获取goroutine
    for {
        h := atomic.LoadAcq(&_p_.runqhead) // load-acquire, synchronize with other consumers
        t := _p_.runqtail
        if t == h {
            return nil, false
        }
        gp := _p_.runq[h%uint32(len(_p_.runq))].ptr()
        if atomic.CasRel(&_p_.runqhead, h, h+1) { // cas-release, commits consume
            return gp, false
        }
    }
}
```

这里首先需要注意的是不管是从runnext还是从循环队列中拿取goroutine都使用了cas操作，这里的cas操作是必需的，因为可能有其他工作线程此时此刻也正在访问这两个成员，从这里偷取可运行的goroutine。

其次，代码中对runqhead的操作使用了`atomic.LoadAcq`和`atomic.CasRel`，它们分别提供了`load-acquire`和`cas-release`语义。

**对于atomic.LoadAcq来说，其语义主要包含如下几条：**

- 原子读取，也就是说不管代码运行在哪种平台，保证在读取过程中不会有其它线程对该变量进行写入；
- 位于`atomic.LoadAcq`之后的代码，对内存的读取和写入必须在`atomic.LoadAcq`读取完成后才能执行，编译器和CPU都不能打乱这个顺序；
- 当前线程执行`atomic.LoadAcq`时可以读取到其它线程最近一次通过`atomic.CasRel`对同一个变量写入的值，与此同时，位于`atomic.LoadAcq`之后的代码，不管读取哪个内存地址中的值，都可以读取到其它线程中位于atomic.CasRel（对同一个变量操作）之前的代码最近一次对内存的写入。

**对于atomic.CasRel来说，其语义主要包含如下几条：**

- 原子的执行比较并交换的操作；
- 位于`atomic.CasRel`之前的代码，对内存的读取和写入必须在`atomic.CasRel`对内存的写入之前完成，编译器和CPU都不能打乱这个顺序；
- 线程执行`atomic.CasRel`完成后其它线程通过`atomic.LoadAcq`读取同一个变量可以读到最新的值，与此同时，位于`atomic.CasRel`之前的代码对内存写入的值，可以被其它线程中位于`atomic.LoadAcq`（对同一个变量操作）之后的代码读取到。

因为可能有多个线程会并发的修改和读取`runqhead`，以及需要依靠runqhead的值来读取runq数组的元素，所以需要使用atomic.LoadAcq和atomic.CasRel来保证上述语义。

我们可能会问，为什么读取p的runqtail成员不需要使用atomic.LoadAcq或atomic.load？因为runqtail不会被其它线程修改，只会被当前工作线程修改，此时没有人修改它，所以也就不需要使用原子相关的操作。

最后，由`p`的`runq`、`runqhead`和`runqtail`这三个成员组成的这个无锁循环队列非常精妙，我们会在后面的章节对这个循环队列进行分析。

## CAS操作与ABA问题

我们知道使用cas操作需要特别注意ABA的问题，那么runqget函数这两个使用cas的地方会不会有问题呢？答案是这两个地方都不会有ABA的问题。原因分析如下：

首先来看对runnext的cas操作。只有跟_p_绑定的当前工作线程才会去修改runnext为一个非0值，其它线程只会把runnext的值从一个非0值修改为0值，然而跟_p_绑定的当前工作线程正在此处执行代码，所以在当前工作线程读取到值A之后，不可能有线程修改其值为B(0)之后再修改回A。

再来看对runq的cas操作。当前工作线程操作的是_p_的本地队列，只有跟_p_绑定在一起的当前工作线程才会因为往该队列里面添加goroutine而去修改runqtail，而其它工作线程不会往该队列里面添加goroutine，也就不会去修改runqtail，它们只会修改runqhead，所以，当我们这个工作线程从runqhead读取到值A之后，其它工作线程也就不可能修改runqhead的值为B之后再第二次把它修改为值A（因为runqtail在这段时间之内不可能被修改，runqhead的值也就无法越过runqtail再回绕到A值），也就是说，代码从逻辑上已经杜绝了引发ABA的条件。

到此，我们已经分析完工作线程从全局运行队列和本地运行队列获取goroutine的代码，由于篇幅的限制，我们下一节再来分析从其它工作线程的运行队列偷取goroutine的流程。






































