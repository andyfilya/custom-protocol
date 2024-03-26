package server

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Addr string
	Port uint16
	Type string

	listener net.Listener
}

func (s *Server) Listen() error {
	for {
		clientConn, err := s.listener.Accept()
		if err != nil {
			log.Println("error in listen server")
			return err
		}

		go handleConn(clientConn)
	}

	return nil
}

func (s *Server) Shutdown() error {
	return s.listener.Close()
}

func InitServer(
	addr, tp string,
	port uint16,
) (*Server, error) {

	l, err := net.Listen(tp, addr+":"+fmt.Sprintf("%d", port))
	if err != nil {
		return nil, err
	}

	return &Server{
		Addr: addr,
		Port: port,

		listener: l,
	}, nil
}
