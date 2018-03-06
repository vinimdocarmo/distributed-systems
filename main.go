package main

import (
	"log"
	"os"

	"./server"
)

func main() {

	s := server.NewServer()

	err := s.Listen("3000")

	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
}
