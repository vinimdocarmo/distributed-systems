package main

import (
	"log"
	"os"

	"./socket"
)

func main() {
	s := socket.NewScoket()

	err := s.Listen("0.0.0.0", "3000")

	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
}
