package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	coon   net.Conn
	server *Server
}

func NewUser(coon net.Conn, server *Server) *User {
	userAddr := coon.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		coon:   coon,
		server: server,
	}

	go user.ListenMessage()

	return user
}

func (t *User) Online() {
	t.server.mapLock.Lock()
	t.server.OnlineUser[t.Name] = t
	t.server.mapLock.Unlock()

	t.server.Broadcast(t, "上线啦")
}

func (t *User) Offline() {
	t.server.mapLock.Lock()
	delete(t.server.OnlineUser, t.Name)
	t.server.mapLock.Unlock()

	t.server.Broadcast(t, "下线啦")
}




func (t *User) ListenMessage() {
	for {
		msg := <-t.C
		t.coon.Write([]byte(msg + "\n"))
	}
}
