package server

import "net"

func handleConn(clientConn net.Conn) {
	defer clientConn.Close()
}
