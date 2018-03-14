package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"./socket"
)

type gui struct {
}

func (g gui) init(mode, addr string) {
	g.clearScreen()

	fmt.Println(`
   _____               __          __ 
  / ___/ ____   _____ / /__ ___   / /_
  \__ \ / __ \ / ___// //_// _ \ / __/
 ___/ // /_/ // /__ / ,<  /  __// /_  
/____/ \____/ \___//_/|_| \___/ \__/  
												
		  `)

	fmt.Printf("🕹️ socket in %s mode\n", mode)

	if mode == "server" {
		serverHost, serverPort, err := net.SplitHostPort(addr)

		if err != nil {
			log.Fatal("⚠️ error reading server address from command line arguments")
			os.Exit(1)
		}

		s := socket.NewServer()

		go s.Listen(serverHost, serverPort)

		fmt.Printf("👂 socket listening on %s\n\n", net.JoinHostPort(serverHost, serverPort))

		for {
			select {
			case ms := <-s.Messages:
				fmt.Printf("📥 %s: %s\n", ms.Addr(), ms.Message())
			}
		}
	} else if mode == "client" {
		peerHost, peerPort, err := net.SplitHostPort(addr)

		if err != nil {
			log.Fatal("⚠️ error reading peer address from command line arguments")
			os.Exit(1)
		}

		c := socket.NewClient()

		go c.Connect(peerHost, peerPort)

		fmt.Printf("🖥️ connected to peer %s\n\n", net.JoinHostPort(peerHost, peerPort))

		for {
			select {
			case connected := <-c.Connected:
				if !connected {
					continue
				}

				reader := bufio.NewReader(os.Stdin)

				for {
					fmt.Print("💬 : ")
					text, _ := reader.ReadString('\n')

					text = strings.Replace(text, "\n", "", -1)

					c.Conn.Write([]byte(text))
				}
			}

		}
	}

}

func (g gui) clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	gui := gui{}

	flagMode := flag.String("mode", "", "start in client or server mode")
	flagAddr := flag.String("address", "", "address to connect to or to listening on")
	flag.Parse()
	mode := strings.ToLower(*flagMode)

	fmt.Print(*flagMode)

	if mode == "server" {
		gui.init(mode, *flagAddr)
	} else if mode == "client" {
		gui.init(mode, *flagAddr)
	}
}
