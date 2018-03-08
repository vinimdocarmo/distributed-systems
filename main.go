package main

import (
	"log"
	"os"

	"./socket"
)

func main() {
	s := socket.NewScoket()

	listenOn := os.Args[1]

	cServer := make(chan int)
	pingTo := os.Args[2]

	go func() {
		err := s.Listen("0.0.0.0", listenOn)

		if err != nil {
			log.Fatal("Error: ", err)
			os.Exit(1)
		}

		cServer <- 0
	}()

	cClient := make(chan int)

	go func() {
		s.Ping("0.0.0.0", pingTo)
		cClient <- 0
	}()

	<-cServer
	<-cClient
}
