package socket

import (
	"bufio"
	"io"
	"net"
	"strings"
)

const (
	crlf = "\r\n"
)

//Socket represents the interface of communication
//between the layers transportion and application
type Socket struct {
	serverAddr string
	peerAddr   string
}

//Listen to on the specified port and set the listener to the server
func (s Socket) Listen(c chan string) {
	l, err := net.Listen("tcp4", s.serverAddr)

	if err != nil {
		c <- "exit"
		return
	}

	c <- "listening"

	for {
		conn, err := l.Accept()

		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			var (
				buf = make([]byte, 1024)
				r   = bufio.NewReader(conn)
				w   = bufio.NewWriter(conn)
			)

			for {
				n, err := r.Read(buf)
				data := string(buf[:n])

				if err == io.EOF {
					conn.Close()
					break
				}

				if err == nil {
					c <- data
					s.Pong(w)
					continue
				}
			}
		}(conn)
	}
}

//Pong responds to client with string "pong"
func (s Socket) Pong(w *bufio.Writer) {
	w.Write([]byte("pong" + crlf))
	w.Flush()
}

//Connect connects to peer in address s.peerAddr
func (s Socket) Connect(c chan string) error {
	conn, err := net.Dial("tcp4", s.peerAddr)

	if err != nil {
		return err
	}

	defer conn.Close()

	c <- "connected"

	return nil
}

//NewSocket creates an instance of Socket
func NewSocket(serverAddr, peerAddr string) Socket {
	return Socket{serverAddr, peerAddr}
}

func isCRLF(data string) bool {
	return strings.HasSuffix(data, crlf)
}
