package main

import "fmt"

/**
Golang 选项模式（Option模式）
*/
func main() {
	message1 := New(1, 123, "message1", "cache1")
	message1.String()
	message2 := NewByOption(WithID(2), WithName("message2"), WithAddress("cache2"), WithPhone(456))
	message2.String()
	message3 := NewByOptionWithoutID(3, WithAddress("cache3"), WithPhone(789), WithName("message3"))
	message3.String()
}

type Message struct {
	id      int
	name    string
	address string
	phone   int
}

func (msg Message) String() {
	fmt.Printf("ID:%d \n- Name:%s6 \n- Address:%s6 \n- phone:%d\n", msg.id, msg.name, msg.address, msg.phone)
}

func New(id, phone int, name, addr string) Message {
	return Message{
		id:      id,
		name:    name,
		address: addr,
		phone:   phone,
	}
}

type Option func(msg *Message)

var DEFAULT_MESSAGE = Message{id: -1, name: "-1", address: "-1", phone: -1}

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

func NewByOption(opts ...Option) Message {
	msg := DEFAULT_MESSAGE
	for _, o := range opts {
		o(&msg)
	}
	return msg
}

func NewByOptionWithoutID(id int, opts ...Option) Message {
	msg := DEFAULT_MESSAGE
	msg.id = id
	for _, o := range opts {
		o(&msg)
	}
	return msg
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
