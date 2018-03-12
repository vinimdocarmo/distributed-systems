package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"./socket"
)

type gui struct {
}

func (g gui) init(c chan string) {
	fmt.Println(`
   _____               __          __ 
  / ___/ ____   _____ / /__ ___   / /_
  \__ \ / __ \ / ___// //_// _ \ / __/
 ___/ // /_/ // /__ / ,<  /  __// /_  
/____/ \____/ \___//_/|_| \___/ \__/  
												
		  `)

	serverHost, serverPort, err := net.SplitHostPort(os.Args[1])

	if err != nil {
		log.Fatal("⚠️ error reading server address from command line arguments")
		os.Exit(1)
	}

	peerHost, peerPort, err := net.SplitHostPort(os.Args[2])

	if err != nil {
		log.Fatal("⚠️ error reading peer address from command line arguments")
		os.Exit(1)
	}

	s := socket.NewSocket(
		net.JoinHostPort(serverHost, serverPort),
		net.JoinHostPort(peerHost, peerPort),
	)

	schan := make(chan string)
	pchan := make(chan string)

	go s.Listen(schan)
	go s.Connect(pchan)

	go func() {
		for {
			str := <-schan

			switch str {
			case "listening":
				fmt.Printf("👂 socket listening on %s\n", net.JoinHostPort(serverHost, serverPort))
			case "exit":
				fmt.Printf("❌ stopped listening on %s\n", net.JoinHostPort(serverHost, serverPort))
				close(schan)
				c <- "exit"
				return
			default:
				fmt.Printf("📥 %s: %s", net.JoinHostPort(peerHost, peerPort), str)
			}
		}
	}()

	go func() {
		for {
			str := <-pchan

			switch str {
			case "connected":
				fmt.Printf("🖥️ connected to peer %s\n", net.JoinHostPort(peerHost, peerPort))
			case "exit":
				fmt.Printf("❌ peer disconnected %s\n", net.JoinHostPort(serverHost, serverPort))
				close(pchan)
				c <- "exit"
			default:
				fmt.Printf("📥 %s: %s", net.JoinHostPort(peerHost, peerPort), str)
			}
		}
	}()

}

func (g gui) chat(serverAddr string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("💬 %s: ", serverAddr)
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		//send text message
	}
}

func main() {
	g := gui{}

	gchan := make(chan string)

	g.init(gchan)

	for {
		str := <-gchan

		if str == "exit" {
			return
		}
	}

}
