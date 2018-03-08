package main

import (
	"log"
	"os"

	"./socket"
)

func main() {
	s := socket.NewScoket()

	listenOn := os.Args[1]
	pingTo := os.Args[2]

	err := s.Listen("0.0.0.0", listenOn)

	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}

	err = s.Ping("0.0.0.0", pingTo)

	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}

}
