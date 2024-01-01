package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	// Uncomment this block to pass the first stage
)

var resolverAddress string

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	flag.StringVar(&resolverAddress, "resolver", "127.0.0.1:2053", "The address of the resolver")
	flag.Parse()

	log.Printf("Forwarding dns server to %s\n", resolverAddress)
	resolver, err := NewResolver(resolverAddress)
	if err != nil {
		log.Fatal(err.Error())
	}

	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:2053")
	if err != nil {
		fmt.Println("Failed to resolve UDP address:", err)
		return
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Failed to bind to address:", err)
		return
	}
	defer udpConn.Close()

	buf := make([]byte, 512)

	for {
		size, source, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}

		// TODO: if resolver flag weren't empty, use local resolver instead.
		ans := NewDNS(buf[:size], resolverAddress)

		// resolver.
		if resolverAddress != "" {
			ans, err = resolver.Serialize(ans)
			if err != nil {
				log.Println(err.Error())
				continue
			}
		}

		_, err = udpConn.WriteToUDP(ans.Bytes(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
