# sync
## 锁`sync.Mutex`
`sync.Mutex`是一个结构体类型
在多协程运行时，一个变量被多个gorouting修改，但是无法保证在这个线程结束前，这个变量没有被其他线程修改，，所以我们可以通过上锁来保证线程结束前不被修改
1. 用法1：如上
```go

type Info struct {

mu sync.Mutex

//

}

func ... (info *Info) {

info.mu.Lock() //上锁

//修改

info.mu.Unlock()//解锁

}

```

2. 用法2：实现一个可上锁的共享缓冲器
	共享缓冲器：多个协程可以共享访问的存储区
	同样也是在结构体变量中引入一个锁变量
	此外，还有一个锁`RWmytex`，通过`Rlock()`允许**多线程可读，单线程可写**，`Lock()`仍是单线程可读写

## 保证被调用函数仅被使用一次的变量类型 `sync.Once`

这是一个变量类型，使用方法如下

```go

var once sync.Once

funcation := func(){

....

}

func main() {

once.Do(funcation)

once.Do(funcation)

once.Do(funcation)//即使有多个语句也只执行一次

}

```

## 等待所有协程完成 `sync.WaitGroup`

这个变量类型所包含的方法可以使主程序等待所有的协程结束

```go

var wg sync.WaitGroup

wg.Add(3)//表示添加了三个协程，不是表示总共3个协程，是添加!

//开始三个协程
defer wg.Done() //使计数减少1
wg.Wait()//等待到计数归0

```

如果协程数量是动态的，有如下方法
动态添加任务
向协程添加一个通道，增加一个协程向协程写入，而原有的每个协程需要从通道中读取才能执行，在新协程中写入之前增加wg计数，实现动态添加

## 条件变量类型 `sync.Cond`

一种结构体类型

使用`sync.NewCond()`函数创建(`func NewCond(l Locker) *Cond`)，同时必须包含一个锁（`Mutex`或`RWMutex`），调用其内的锁可用`cond.L.Lock()`

```go

var mu sync.Mutex

cond := sync.NewCond(&mu)

//使用RWMutex时导入要用方法xxx.RLocker()在NewCond的括号内

```

方法如下
1. 等待 `Wait()`
	`cond.Wait()`
	会自动释放锁，并挂起当前gorouting开始等待通知
	必须使用循环检查，防止虚假唤醒
2. 通知一个等待协程 Signal()
	通知等待时间最长的
	`cond.Signal()`
3. 通知所有等待协程 Broadcast()
	`cond.Broadcast()`

# sync/atomic
