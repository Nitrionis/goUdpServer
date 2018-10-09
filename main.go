package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"net"
)

var clients *list.List

func handleClient(socket net.Conn) {
	defer log.Println("Client disconected.")
	for {
		buffer, err := bufio.NewReader(socket).ReadString('\n')
		if err != nil {
			log.Fatal(err)
			socket.Close()
			return
		}
		fmt.Println(buffer)
		for i:= clients.Front(); i != nil; i = i.Next() {
			writer := bufio.NewWriter(i.Value.(net.Conn))
			writer.WriteString(buffer)
			writer.WriteString(string(0x000A))
			writer.Flush()
		}
	}
}

func main() {
	fmt.Println("Server start!")

	clients = list.New()

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		client, err := server.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("New user accept %s", client.RemoteAddr())
		clients.PushBack(client)
		go handleClient(client)
	}

	fmt.Println("Server closed!")
}