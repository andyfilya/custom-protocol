package main

import (
	"fmt"
	"github.com/andyfilya/customprotocol/internal/rules"
	"github.com/andyfilya/customprotocol/pkg/server"
	"log"
	"net"
)

const (
	addr        = "localhost"
	tp          = "tcp"
	port uint16 = 8081
)

func main() {
	log.Println("server init")
	serv := server.InitServer(addr, tp, port)
	log.Println("successful init server")
	serv.Listen(func(conn net.Conn) {
		log.Println("in handler function")
		var res rules.Message

		rules.ReadJSON(conn, &res)

		fmt.Printf("at server: %v \n \n", res)
		msg := rules.Message{
			Header: "from server",
			Body:   "hello",
		}

		log.Printf("server send msg : %v", msg)
		rules.SendJSON(conn, msg)
	})
}
