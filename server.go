package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (server *Server) Handler(accept net.Conn) {
	fmt.Println("handler accept")
	fmt.Println(accept)
}

func (server *Server) start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("net.listen err")
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {

		}
	}(listen)

	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err")
		}
		server.Handler(accept)
	}

}
