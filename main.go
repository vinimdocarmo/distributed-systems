package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type gui struct {
}

func (g gui) init() {
	fmt.Println(`
   _____               __          __ 
  / ___/ ____   _____ / /__ ___   / /_
  \__ \ / __ \ / ___// //_// _ \ / __/
 ___/ // /_/ // /__ / ,<  /  __// /_  
/____/ \____/ \___//_/|_| \___/ \__/  
												
		  `)
}

func main() {
	// s := socket.NewScoket()
	g := gui{}

	serverHost, serverPort, err := net.SplitHostPort(os.Args[1])

	if err != nil {
		log.Fatal("Error reading address from command line arguments")
		os.Exit(1)
	}

	peerHost, peerPort, err := net.SplitHostPort(os.Args[2])

	if err != nil {
		log.Fatal("Error reading address from command line arguments")
		os.Exit(1)
	}

	g.init()
	fmt.Printf("👂 socket listening on \t%s\n", net.JoinHostPort(serverHost, serverPort))
	fmt.Printf("🖥️ connected to peer \t%s\n\n\n", net.JoinHostPort(peerHost, peerPort))

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("💬 %s: ", net.JoinHostPort(serverHost, serverPort))
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Printf("📥 %s: \n", net.JoinHostPort(peerHost, peerPort))
		}
	}

}
