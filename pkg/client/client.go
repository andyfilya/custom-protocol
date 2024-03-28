package client

import (
	"fmt"
	"github.com/andyfilya/customprotocol/internal/rules"
	"log"
	"net"
)

type Client struct{}

func (cli *Client) Connect(tp string, addr string, handler rules.Handler) error {
	conn, err := net.Dial(tp, addr)
	if err != nil {
		log.Printf("error to connect dial %s to addr : %s", tp, addr)
		return err
	}
	fmt.Println("client ok")
	defer conn.Close()
	handler(conn)
	return nil
}
