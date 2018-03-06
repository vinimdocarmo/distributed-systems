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

		go handlerPing(conn)
	}
}

func handlerPing(conn net.Conn) {

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
			if isPing(data) {
				log.Print("received: ", data)
				continue ILOOP
			}
			if isCRLF(data) {
				break ILOOP
			}

		default:
			log.Fatalf("receive data failed: %s", err)
			return
		}

	}
	pong(w)
	log.Printf("send: %s", "pong")
}

func pong(w *bufio.Writer) {
	w.Write([]byte("pong" + CRLF))
	w.Flush()
}

func isPing(data string) bool {
	return strings.HasPrefix(data, "ping")
}

func isCRLF(data string) bool {
	return strings.HasSuffix(data, CRLF)
}
