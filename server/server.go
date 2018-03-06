package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

const (
	CRLF = "\r\n"
)

//Server represents a server in a communication client-server
type Server struct {
	listener net.Listener
}

//NewServer creates an instance of Server
func NewServer() Server {
	return Server{}
}

//Pong responds to client with string "PONG"
func (s Server) Pong() {

}

//Listen to on the specified port and set the listener to the server
func (s Server) Listen(port string) error {
	address := "0.0.0.0:" + port
	l, err := net.Listen("tcp4", address)

	if err != nil {
		return err
	}

	log.Printf("socket listening on %s\n", address)

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

	defer conn.Close()

ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Received:", data)
			if isCRLF(data) {
				break ILOOP
			}

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}

	}
	pong(w)
	log.Printf("Send: %s", "PONG")
}

func pong(w *bufio.Writer) {
	w.Write([]byte("PONG" + CRLF))
	w.Flush()
}

func isCRLF(data string) bool {
	return strings.HasSuffix(data, CRLF)
}
