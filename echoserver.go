package echoserver

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) Run() error {
	server, err := net.Listen("tcp", s.Host+":"+strconv.Itoa(s.Port))
	if err != nil {
		return err
	}
	conns := s.ClientConns(server)
	for {
		go s.HandleConn(<-conns)
	}
	return nil
}

func (s *Server) ClientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				fmt.Printf("couldn't accept: " + err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func (s *Server) HandleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil { // EOF, or worse
			break
		}
		client.Write(line)
	}
}
