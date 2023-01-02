package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	lock      sync.RWMutex
	Message   chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	//发送消息
	this.Message <- sendMsg
}

func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.lock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.lock.Unlock()
	}
}

func (this *Server) Handle(con net.Conn) {
	user := NewUser(con)
	//加锁
	this.lock.Lock()
	this.OnlineMap[user.Name] = user
	this.lock.Unlock()

	this.BroadCast(user, "已上线")

	//客户端是否存活channel
	isLive := make(chan bool)

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		//for 循环很重要
		for {
			n, err := con.Read(buf)
			if n == 0 {
				delete(this.OnlineMap, user.Name)
				//关闭链接
				this.BroadCast(user, "下线")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Read error:", err)
				return
			}
			isLive <- true
			//去除 \n
			msg := string(buf[:n-1])
			if msg == "who" {
				//查询在线用户
				this.lock.Lock()
				for _, onlineUser := range this.OnlineMap {
					sendMsg := "[" + onlineUser.Addr + "]" + onlineUser.Name + ": 在线"
					user.C <- sendMsg
				}
				this.lock.Unlock()
			} else if len(msg) > 7 && msg[:6] == "rename" {
				//rename|张三
				newName := strings.Split(msg, "|")[1]
				//更改用户名称
				this.lock.Lock()
				if this.OnlineMap[newName] != nil {
					user.C <- "用户名已存在！"
					return
				}
				delete(this.OnlineMap, user.Name)
				user.Name = newName
				this.OnlineMap[newName] = user
				this.lock.Unlock()
			} else if len(msg) > 4 && msg[:3] == "to|" {
				//私聊格式 to|name|msg
				splitMsg := strings.Split(msg, "|")
				toName := splitMsg[1]
				toMsg := splitMsg[2]
				//1.获取该用户
				toUser := this.OnlineMap[toName]
				if toUser == nil {
					user.SendMessage("发送用户不存在！")
					return
				}
				if len(toMsg) == 0 {
					user.SendMessage("消息不能为空！")
					return
				}
				//2.发送消息
				toUser.C <- "[" + user.Name + "] 对您说：" + toMsg
			} else {
				this.BroadCast(user, msg)
			}
		}
	}()

	//当前handler阻塞
	//select {}
	for {
		select {
		case <-isLive:
			//do nothing, update timer
		case <-time.After(300 * time.Second):
			//0.发送消息
			user.SendMessage("你已超时被强制下线！")
			//1.关闭用户连接
			user.conn.Close()
			////2.关闭chan
			close(user.C)
			////3.删除map
			this.lock.Lock()
			delete(this.OnlineMap, user.Name)
			this.lock.Unlock()
			//4 return
			return
		}
	}
	//fmt.Println(user.Name+ " 链接成功")
}

func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("listen err: ", err)
		return
	}
	//关闭socket
	defer listen.Close()
	//监听
	go this.ListenMessage()
	for {
		//获取链接
		con, err := listen.Accept()
		if err != err {
			fmt.Println("listen accept err: ", err)
			continue
		}
		//处理链接
		go this.Handle(con)
	}

}
