package main

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	header := &Header{
		ID:    1234,
		QR:    true,
		RCODE: 2,
	}

	newheader := new(Header)

	newheader.Parse(header.ToBytes())

	header.RCODE = 4
	assert.Equal(t, header, newheader)
}

func TestQuestion(t *testing.T) {
	q := &Question{
		Name:  "codecrafters.io",
		Type:  1,
		Class: 1,
	}

	fmt.Println(string(q.ToBytes()))
}

func createTestDNSPacket() []byte {
	var buf bytes.Buffer

	// Header
	buf.Write([]byte{0x00, 0x01}) // ID
	buf.Write([]byte{0x01, 0x00}) // Flags
	buf.Write([]byte{0x00, 0x02}) // QDCOUNT
	buf.Write([]byte{0x00, 0x00}) // ANCOUNT
	buf.Write([]byte{0x00, 0x00}) // NSCOUNT
	buf.Write([]byte{0x00, 0x00}) // ARCOUNT

	// Question 1: abc.longassdomainname.com
	buf.Write([]byte{0x03}) // Length of "abc"
	buf.WriteString("abc")
	buf.Write([]byte{0x11}) // Length of "longassdomainname"
	buf.WriteString("longassdomainname")
	buf.Write([]byte{0x03}) // Length of "com"
	buf.WriteString("com")
	buf.Write([]byte{0x00})       // Null terminator
	buf.Write([]byte{0x00, 0x01}) // QTYPE
	buf.Write([]byte{0x00, 0x01}) // QCLASS

	// Question 2: def.longassdomainname.com
	buf.Write([]byte{0x03}) // Length of "def"
	buf.WriteString("def")
	buf.Write([]byte{0x11}) // Length of "longassdomainname"
	buf.WriteString("longassdomainname")
	buf.Write([]byte{0x03}) // Length of "com"
	buf.WriteString("com")
	buf.Write([]byte{0x00})       // Null terminator
	buf.Write([]byte{0x00, 0x01}) // QTYPE
	buf.Write([]byte{0x00, 0x01}) // QCLASS

	return buf.Bytes()
}

func TestCompressionMessageQuestion(t *testing.T) {
	question := createTestDNSPacket()

	dns := NewDNS(question, "8.8.8.8:53")

	log.Printf("dns: %+v\n", dns)

	for _, v := range dns.Question {
		log.Printf("question: %+v\n", v)
	}
}
