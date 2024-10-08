package main

import (
	"fmt"
	"sync"
)

func main() {
	//waitGroup()
	once()
}

func waitGroup() {
	var wg = sync.WaitGroup{}
	var N int = 10
	wg.Add(N)
	for i := 0; i < N; i++ {

		go func(id int) {
			fmt.Println(id)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
}

func once() {
	var once = sync.Once{}
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(id int) {
			once.Do(func() {
				fmt.Println(id)
			})
			done <- struct{}{}
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}

func exampleCond() {

	// 初始化一个互斥锁
	// 可以创建为其他结构体的字段，零值为解锁状态
	// Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁
	var locker = new(sync.Mutex)

	// 使用锁locker创建一个*sync.Cond
	// Cond实现了一个条件变量，一个线程集合地，供线程等待或者宣布某事件的发生
	c := sync.NewCond(locker)

	// 锁定
	c.L.Lock()

	// 唤醒所有等待c的线程。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定
	c.Broadcast()

	// 唤醒等待c的一个线程（如果存在）。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定
	c.Signal()

	// 解锁
	c.L.Unlock()

	// 自行解锁c.L并阻塞当前线程，在之后线程恢复执行时，Wait方法会在返回前锁定c.L
	// 和其他系统不同，Wait除非被Broadcast或者Signal唤醒，不会主动返回
	//c.Wait()
}

func examplePool() {

	// 初始化一个可以分别存取的临时对象的集合
	// Pool中保存的任何item都可能随时不做通告的释放掉。如果Pool持有该对象的唯一引用，这个item就可能被回收
	// Pool可以安全的被多个线程同时使用
	// Pool的目的是缓存申请但未使用的item用于之后的重用，以减轻GC的压力。也就是说，让创建高效而线程安全的空闲列表更容易。但Pool并不适用于所有空闲列表
	// Pool的合理用法是用于管理一组静静的被多个独立并发线程共享并可能重用的临时item。Pool提供了让多个线程分摊内存申请消耗的方法
	// 管理着短寿命对象的空闲列表不适合使用Pool，因为这种情况下内存申请消耗不能很好的分配。这时应该由这些对象自己实现空闲列表
	var pool = new(sync.Pool)

	// (可选)指定当get返回nil时生成值的函数。它不能与get调用同时更改
	pool.New = func() interface{} {
		return 5
	}

	// 将x放入池中
	pool.Put(10)

	// 从池中选择任意一个item，删除其在池中的引用计数，并提供给调用者
	// Get方法也可能选择无视内存池，将其当作空的。调用者不应认为Get的返回这和传递给Put的值之间有任何关系
	item := pool.Get()
	fmt.Println(item)
}

func exampleRWMutex() {

	// 初始化一个读写互斥锁
	// 该锁可以被同时多个读取者持有或唯一个写入者持有
	// RWMutex可以创建为其他结构体的字段；零值为解锁状态
	// RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁
	var rw = new(sync.RWMutex)

	// 将rw锁定为写入状态，禁止其他线程读取或者写入
	rw.Lock()

	// 解除rw的写入锁状态，如果未加写入锁会导致运行时错误
	rw.Unlock()

	// 将rw锁定为读取状态，禁止其他线程写入，但不禁止读取
	rw.RLock()

	// 解除rw的读取锁状态，如果m未加读取锁会导致运行时错误
	rw.RUnlock()

	// 返回一个互斥锁，通过调用rw.Rlock和rw.Runlock实现了Locker接口
	rw.RLocker()
}

func exampleMap() {

	// 声明并发Map
	var m sync.Map

	// 将键值对保存到sync.Map
	m.Store("a", "Hello World!")
	m.Store("b", "Hello Gopher!")
	m.Store("c", "Hello zc!")

	// 从sync.Map中根据键取值
	fmt.Println(m.Load("c"))

	// 根据键删除对应的键值对
	m.Delete("c")

	// 遍历所有sync.Map中的键值对
	m.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}