package main

import (
	"github.com/andyfilya/customprotocol/pkg/server"
	"log"
)

const (
	addr        = "localhost"
	tp          = "tcp"
	port uint16 = 8081
)

func main() {
	serv, err := server.InitServer(addr, tp, port)
	if err != nil {
		log.Fatalf("error init server with port %d, [%v]", port, err)
	}
	defer func() {
		err := serv.Shutdown()
		if err != nil {
			log.Fatalf("error shutdown srv [%v]", err)
		}
	}()
	serv.Listen()
}
