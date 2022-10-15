// 创建型模式（Creational Pattern）之
// 01:单例模式（Singleton Pattern）
// 介绍：一个类仅有一个实例，并提供一个访问它的全局访问点。
// 应用：线程池、全局缓存、对象池(被建模成单例的对象都有“中心点”的含义，比如线程池就是管理所有线程的中心)
// 实现：（1）限制调用者直接实例化该对象；（2）为该对象的单例提供一个全局唯一的访问方法。
//       c++：类的构造函数设计成私有的，并提供一个static方法去访问该类的唯一实例
// 		 go：单例结构体设计成首字母小写， 实现一个首字母大写的访问函数
// 饿汉模式：实例在系统加载的时候就已经完成了初始化
// 懒汉模式：等到对象被使用的时才会去初始化它，从而一定程度上节省了内存但，带来线程安全问题
// 可以使用普通加锁，或者更高效的双重检验锁来优化
// golang 使用sync.Once

package msgpool

import "sync"

// 消息池
type messagePool struct {
	pool *sync.Pool
}

type Message struct {
	Count int
}

// golang singleton pattern example1: 饿汉模式
/****************************************************
// 消息池单例
var msgPool = &messagePool{
	pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
}

// 访问消息池的唯一方法
func Instance() *messagePool {
	return msgPool
}

// 添加消息
func (m *messagePool) AddMsg(msg *Message) {
	m.pool.Put(msg)
}

// 获取消息
func (m *messagePool) GetMsg() *Message {
	return m.pool.Get().(*Message)
}
******************************************************/
// 单例模式的“懒汉模式”实现
var once = &sync.Once{}

// 消息池单例，在首次调用时初始化
var msgPool *messagePool

// 全局唯一获取消息池pool到方法
func Instance() *messagePool {
	// 在匿名函数中实现初始化逻辑，Go语言保证只会调用一次
	once.Do(func() {
		msgPool = &messagePool{
			// 如果消息池里没有消息，则新建一个Count值为0的Message实例
			pool: &sync.Pool{New: func() interface{} { return &Message{Count: 0} }},
		}
	})
	return msgPool
}
