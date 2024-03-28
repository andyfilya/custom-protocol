package main

import (
	"fmt"
	"github.com/andyfilya/customprotocol/internal/rules"
	"github.com/andyfilya/customprotocol/pkg/client"
	"log"
	"net"
)

func main() {
	cli := client.Client{}
	cli.Connect("tcp", "localhost:8081", func(conn net.Conn) {
		defer conn.Close()
		msg := rules.Message{
			Header: "from client",
			Body:   "hello",
		}

		log.Printf("client send msg : %v", msg)

		rules.SendJSON(conn, msg)

		var res rules.Message

		rules.ReadJSON(conn, &res)

		fmt.Printf("in client : %v", res)
	})
}
