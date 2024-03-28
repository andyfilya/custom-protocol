package server

import (
	"fmt"
	"github.com/andyfilya/customprotocol/internal/rules"
	"log"
	"net"
	"time"
)

type Server struct {
	Addr string
	Port uint16
	Type string
}

// Listen function start process of listening tcp port
// handler rules.Handler is helper function for handle message in chan
func (s *Server) Listen(handler rules.Handler) error {
	srv, err := net.Listen(s.Type, s.Addr+":"+fmt.Sprintf("%d", s.Port))
	if err != nil {
		log.Println("error to listen srv")
		return err
	}
	defer srv.Close()

	ch := make(chan net.Conn)
	ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()

	// start goroutine which accept connection from clients
	go func() {
		clientConn, err := srv.Accept()
		if err != nil {
			log.Println("error in listen server")
			return
		}

		ch <- clientConn // send connection to channel
	}()

	// unlimited cycle with select
	for {
		select {
		// if in ch have connection -> start goroutine which read message from client
		case clientConn := <-ch:
			fmt.Println("connected")
			time.Sleep(1 * time.Second)
			go func(clientConn net.Conn) {
				handler(clientConn)

				defer clientConn.Close() // make sure that client connections is closed
			}(clientConn)
		// if time 1s -> say you wait for client
		case time := <-ticker.C:
			fmt.Printf("wait for connection %v \n", time)
		}
	}
	return nil
}

// InitServer initialize server with propagated host, port and type of connection
func InitServer(
	addr, tp string,
	port uint16,
) *Server {

	return &Server{
		Addr: addr,
		Port: port,
		Type: tp,
	}
}
