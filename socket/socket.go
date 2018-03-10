package socket

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

const (
	crlf = "\r\n\r\n"
)

//Socket represents the interface of communication
//between the layers transportion and application
type Socket struct {
	listener net.Listener
}

//Listen to on the specified port and set the listener to the server
func (s Socket) Listen(addr, port string) error {
	address := addr + ":" + port
	l, err := net.Listen("tcp4", address)

	if err != nil {
		return err
	}

	log.Printf("[server] listening on %s\n", address)

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatalln(err)
			continue
		}

		go func() {
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
					if isCRLF(data) {
						log.Print("[server] received: ", data)
						s.Pong(w)
						log.Printf("[server] response: %s", "pong")

						break ILOOP
					}
				default:
					log.Fatalf("[server] receive data failed: %s", err)
					return
				}

			}
		}()
	}
}

//Pong responds to client with string "pong"
func (s Socket) Pong(w *bufio.Writer) {
	w.Write([]byte("pong" + crlf))
	w.Flush()
}

//Ping dials to server with string "ping"
func (s Socket) Ping(host, port string) {

	for {
		conn, err := net.Dial("tcp4", net.JoinHostPort(host, port))

		if err != nil {
			log.Println("[client] couldn't dial to " + host + ":" + port)
			time.Sleep(2 * time.Second)
			continue
		}

		conn.Write([]byte("ping"))
		conn.Write([]byte("\r\n\r\n"))

		log.Printf("[client] request: %s", "ping")

		conn.Close()
		time.Sleep(2 * time.Second)
	}
}

//NewScoket creates an instance of Socket
func NewScoket() Socket {
	return Socket{}
}

func isCRLF(data string) bool {
	return strings.HasSuffix(data, crlf)
}
