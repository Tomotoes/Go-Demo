package main

import (
	"flag"
	"fmt"
)

type Server struct {
	Host string
	Port int
}

func NewServer(opts ...func(server *Server)) *Server {
	s := &Server{}
	for _, op := range opts {
		op(s)
	}
	return s
}

func Host(host string) func(*Server) {
	return func(server *Server) {
		server.Host = host
	}
}

func Port(port int) func(*Server) {
	return func(server *Server) {
		server.Port = port
	}
}

func main() {
	var host = flag.String("host", "127.0.0.1", "host")
	var port = flag.Int("port", 1080, "port")
	s := NewServer(
		Host(*host),
		Port(*port),
	)
	fmt.Printf("server host:%s,port: %d\n", s.Host, s.Portk)
}
