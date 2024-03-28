package rules

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Handler func(conn net.Conn)

// PROTOCOL LOOK LIKE THIS
// LENh| HEADER | LENb | BODY -> 2 bytes (length of header + body)

// Message is struct how look like message in custom protocol
type Message struct {
	Header string
	Body   string
}

func Read(conn net.Conn) []byte {
	headerL := make([]byte, 2) // read the length of header message
	h, _ := conn.Read(headerL)
	hl := binary.BigEndian.Uint16(headerL[:h])
	header := make([]byte, hl)
	conn.Read(header)
	bodyL := make([]byte, 2) // read the length of body message
	b, _ := conn.Read(bodyL)
	bl := binary.BigEndian.Uint16(bodyL[:b])
	body := make([]byte, bl)
	conn.Read(body)

	msg := Message{
		Header: string(header),
		Body:   string(body),
	}

	bytes, _ := json.Marshal(msg)

	return bytes
}

// Write function need to write the message by the rules of protocol
func Write(m []byte, conn net.Conn) {
	log.Println("in write function helper")
	_, err := conn.Write(m) // incoming struct with fields have type uint16
	if err != nil {
		log.Println("error write with net package")
		return
	}
	log.Println("successful write to connection with net package")
}

func ReadJSON(conn net.Conn, v any) error {
	log.Println("in read function")
	data := Read(conn)
	fmt.Printf("data : %v \n\n", data)
	return json.Unmarshal(data, v)
}

func SendJSON(conn net.Conn, msg Message) error {
	log.Println("started send json to server")
	var res []byte // the sequence of bytes will send to server or client

	headerLength := uint16(len(msg.Header)) // length of header
	bodyLength := uint16(len(msg.Body))     // length of body

	hL := make([]byte, 2) // slice with header length
	binary.BigEndian.PutUint16(hL, headerLength)
	res = append(res, hL...) // append the len of header in uint16
	header, err := json.Marshal(msg.Header)
	if err != nil {
		log.Println("error to write msg in connection")
		return err
	}
	res = append(res, header...) // append the header

	bL := make([]byte, 2)
	binary.BigEndian.PutUint16(bL, bodyLength)
	res = append(res, bL...) // append the len of body in uint16

	body, err := json.Marshal(msg.Body)
	if err != nil {
		log.Println("error to write message in connection")
		return err
	}
	res = append(res, body...) // append the body of message

	Write(res, conn)
	log.Println("successful write msg in connection")
	return nil
}
