package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/goTalk2/client_app/client"
	"github.com/goTalk2/client_app/server"
)

var (
	msgc       = make(chan string)
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func input() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		msgc <- text
	}
}

func main() {
	fmt.Println("start the program")
	waitc := make(chan struct{})
	client.InitChatClient(serverAddr)

	// start the server thread
	go func() {
		server.InitChatServer()
		close(waitc)
	}()

	// start the client thread
	go func() {
		for {
			msg := <-msgc // a message to send
			client.Chat(msg)
		}
		close(waitc)
	}()

	// start the input thread
	go input()

	<-waitc
}