package server

import (
	"bufio"
)

//Server represents a server in a communication client-server
type Server interface {
	Listen(addr, port string) error
	Pong(w *bufio.Writer)
}
