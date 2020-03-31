package main

import "fmt"

//参考文章：https://blog.csdn.net/Lazyboy_/article/details/103289750

//一种优雅的方式来解决参数传递问题，不用初始化属性值的时候都去修改new函数，特别适合结构体有些默认的值，如果不填就取默认值，如文中的DefaultMessage
//同时也给了调用方自己控制的逻辑，可以选择覆盖或者修改，使用WithXXX()方法

/*
Option模式的优缺点
优点
	支持传递多个参数，并且在参数个数、类型发生变化时保持兼容性
	任意顺序传递参数
	支持默认值
	方便拓展
缺点
	增加许多function，成本增大
	参数不太复杂时，尽量少用
*/
type Message struct {
	id      int
	name    string
	address string
	phone   int
}

func (msg Message) String() {
	fmt.Printf("ID:%d \n- Name:%s \n- Address:%s \n- phone:%d\n", msg.id, msg.name, msg.address, msg.phone)
}

//最传统的方式
func New(id, phone int, name, addr string) Message {
	return Message{
		id:      id,
		name:    name,
		address: addr,
		phone:   phone,
	}
}

type Option func(msg *Message)

var DefaultMessage = Message{id: -1, name: "-1", address: "-1", phone: -1}

func WithID(id int) Option {
	return func(m *Message) {
		m.id = id
	}
}

func WithName(name string) Option {
	return func(m *Message) {
		m.name = name
	}
}

func WithAddress(addr string) Option {
	return func(m *Message) {
		m.address = addr
	}
}

func WithPhone(phone int) Option {
	return func(m *Message) {
		m.phone = phone
	}
}

//根据Option来修改默认值
func NewByOption(opts ...Option) Message {
	msg := DefaultMessage
	for _, o := range opts {
		o(&msg)
	}
	return msg
}

//id取传入的默认值
func NewByOptionWithoutID(id int, opts ...Option) Message {
	msg := DefaultMessage
	msg.id = id
	for _, o := range opts {
		o(&msg)
	}
	return msg
}

func main() {

	message1 := New(1, 123, "message1", "cache1")
	message1.String()
	message2 := NewByOption(WithID(2), WithName("message2"), WithAddress("cache2"), WithPhone(456))
	message2.String()
	message3 := NewByOptionWithoutID(3, WithAddress("cache3"), WithPhone(789), WithName("message3"))
	message3.String()
}

/*
Output
ID:1
- Name:message1
- Address:cache1
- phone:123
ID:2
- Name:message2
- Address:cache2
- phone:456
ID:3
- Name:message3
- Address:cache3
- phone:789
*/
