package socket

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type socketMessage struct {
	m    string
	conn net.Conn
}

func (sm socketMessage) Message() string {
	return sm.m
}

func (sm socketMessage) Addr() string {
	return sm.conn.RemoteAddr().String()
}

//Server represents a server on client-server architecture
type Server struct {
	Messages     chan socketMessage
	Connected    chan net.Conn
	Disconnected chan net.Conn
}

//NewServer creates a new instance of Server
func NewServer() Server {
	return Server{
		Messages:     make(chan socketMessage),
		Connected:    make(chan net.Conn),
		Disconnected: make(chan net.Conn),
	}
}

//Listen announces on the local network address.
func (s Server) Listen(host, port string) {
	addr := net.JoinHostPort(host, port)
	l, err := net.Listen("tcp", addr)

	if err != nil {
		panic(fmt.Sprintf("error listening on %s", addr))
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Printf("error accepting connection from %v", conn.RemoteAddr().String())
		}
		go handleConnection(s, conn)
	}
}

func handleConnection(s Server, conn net.Conn) {
	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	s.Connected <- conn

	defer conn.Close()

	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		if err == io.EOF {
			s.Disconnected <- conn
			break
		}

		if err == nil {
			s.Messages <- socketMessage{
				m:    data,
				conn: conn,
			}
			w.Write([]byte("200 OK"))
			w.Flush()
			continue
		}
	}
}
