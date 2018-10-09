package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleClient(conn net.Conn) {
	defer log.Println("Client disconected.")
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	for {
		buffer, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			conn.Close()
			return
		}
		fmt.Println(buffer)

		w.WriteString(buffer)
		w.Flush()
	}
}

func main() {
	fmt.Println("Server start!")

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Printf("New user accept %s", conn.RemoteAddr())
		go handleClient(conn)
	}

	fmt.Println("Server closed!")
}
