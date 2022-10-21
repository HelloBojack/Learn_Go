package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	OnlineUser map[string]*User
	mapLock    sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:         ip,
		Port:       port,
		OnlineUser: make(map[string]*User),
		Message:    make(chan string),
	}
	return server
}

func (t *Server) ListenMessage() {
	for {
		msg := <-t.Message

		t.mapLock.Lock()
		for _, user := range t.OnlineUser {
			user.C <- msg
		}
		t.mapLock.Unlock()
	}
}

func (t *Server) Broadcast(user *User, msg string) {
	if msg == "who" {
		t.mapLock.Lock()
		for _, usr := range t.OnlineUser {
			sendMsg := "[" + usr.Addr + "]" + usr.Name + ":" + "在线ing..."
			user.C <- sendMsg
		}
		t.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := msg[7:]
		user.ChangName(newName)
	} else if msg[:3] == "to|" {
		name := strings.Split(msg, "|")[1]
		cnt := strings.Split(msg, "|")[2]
		usr, ok := t.OnlineUser[name]
		if ok {
			usr.C <- user.Name + "对您说：" + cnt
		}
	} else {
		sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
		t.Message <- sendMsg
	}
}

func (t *Server) Handler(conn net.Conn) {
	user := NewUser(conn, t)

	user.Online()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)

			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("error")
				return
			}

			msg := string(buf[:n-1])

			t.Broadcast(user, msg)

		}
	}()

}

func (t *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", t.Ip, t.Port))
	if err != nil {
		fmt.Printf("net.Listen error:%s", err)
		return
	}
	defer listener.Close()

	fmt.Println("Start")
	go t.ListenMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener.Accept error:%s", err)
			continue
		}
		go t.Handler(conn)
	}
}
