package main

import (
	"log"
	"net"
	"os"

	"./socket"
)

func main() {
	s := socket.NewScoket()

	host1, port1, err := net.SplitHostPort(os.Args[1])

	if err != nil {
		log.Fatal("Error reading address from command line arguments")
		os.Exit(1)
	}

	host2, port2, err := net.SplitHostPort(os.Args[2])

	if err != nil {
		log.Fatal("Error reading address from command line arguments")
		os.Exit(1)
	}

	cServer := make(chan int)

	go func() {
		err := s.Listen(host1, port1)

		if err != nil {
			log.Fatal("Error: ", err)
			os.Exit(1)
		}

		cServer <- 0
	}()

	cClient := make(chan int)

	go func() {
		s.Ping(host2, port2)
		cClient <- 0
	}()

	<-cServer
	<-cClient
}
