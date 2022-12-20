package main

import (
	"fmt"
	"net"
)

type User struct {
	Name    string
	Address string
	C       chan string
	conn    net.Conn
}

func NewUser(conn net.Conn) *User {

	userAddress := conn.RemoteAddr().String()
	user := &User{
		Name:    userAddress,
		Address: userAddress,
		C:       make(chan string),
		conn:    conn,
	}
	return user
}

func (user *User) ListenMessage() {
	for {
		msg := <-user.C
		_, err := user.conn.Write([]byte(msg + "\n"))
		if err != nil {
			fmt.Println("user.conn err", err)
		}

	}
}
