package client

//Client represents a client in a communication client-server
type Client interface {
	Ping(addr, port string) error
}
