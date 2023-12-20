package main

import (
	"fmt"
	"net"
	// Uncomment this block to pass the first stage
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

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

		receivedData := string(buf[:size])
		fmt.Printf("Received %d bytes from %s: %s\n", size, source, receivedData)

		responseheader := &Header{
			ID:      1234,
			QR:      true,
			QDCOUNT: 1,
			ANCOUNT: 1,
		}

		question := &Question{
			Name:  "codecrafters.io",
			Type:  int(TypeNameToValue("A")),
			Class: int(ClassNameToValue("IN")),
		}

		answer := &Answer{
			Name:  "codecrafters.io",
			Type:  int(TypeNameToValue("A")),
			Class: int(ClassNameToValue("IN")),
			TTL:   60,
			Data:  "8.8.8.8",
		}

		response := &DNS{
			Header:   responseheader,
			Question: question,
			Answer:   answer,
		}

		_, err = udpConn.WriteToUDP(response.Bytes(), source)
		if err != nil {
			fmt.Println("Failed to send response:", err)
		}
	}
}
