package socket

import (
	"fmt"
	"net"
)

//Client represents a server on client-server architecture
type Client struct {
	Conn      net.Conn
	Connected chan bool
}

//NewClient creates a new instance of Client
func NewClient() Client {
	return Client{Connected: make(chan bool)}
}

//Connect connects to the host on port
func (c *Client) Connect(host, port string) {
	addr := net.JoinHostPort(host, port)

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		panic(fmt.Sprintf("error dialing to %s", addr))
	}

	c.Conn = conn

	c.Connected <- true
}
