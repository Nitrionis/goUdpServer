package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
		/*go*/ serve(pc, addr, buf[:n])
		time.Sleep(100 * time.Millisecond)
	}
}

func serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	fmt.Println(string(buf))
	pc.WriteTo(buf, addr)
}
