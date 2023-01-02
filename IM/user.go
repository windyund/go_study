package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

func NewUser(con net.Conn) *User {
	userAddr := con.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: con,
	}
	//开启goroutine 监听chan,如果有消息，就通过con发送客户端
	go user.ListenMessage()
	return user
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

func (this *User) SendMessage(msg string) {
	this.conn.Write([]byte(msg + "\n"))
}
