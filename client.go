package main

import (
	"net"
	"bufio"
	"fmt"
)

type Client struct {
	incoming chan string // maze response
	outgoing chan string // request
	conn net.Conn
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func (client *Client) Read() {
	for {
		line, _ := client.reader.ReadString('\n')
		fmt.Println(line)
		client.incoming <- line
		fmt.Printf("Got an incoming message %s. \n", line)
	}
}

func (client *Client) Write() {
	for data := range client.outgoing {
		fmt.Println("going to write")
		data = data + "\n"
		client.conn.Write([]byte(data))
	}
}

func (client *Client) Listen() {
	go client.Read()
	go client.Write()
}
 
func NewClient(connection net.Conn) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client{
		incoming: make(chan string),
		outgoing: make(chan string),
		conn: connection,
		reader: reader,
		writer: writer,
	}

	client.Listen()

	return client
}