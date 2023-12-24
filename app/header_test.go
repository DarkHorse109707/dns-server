package main

import (
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

func TestCompressionMessageQuestion(t *testing.T) {
	// question := []byte{
	// 	0x07, 'e', 'x', 'a', 'm', 'p', 'l', 'e',
	// 	0x03, 'c', 'o', 'm',
	// 	0xC0, 0x0C, // Compression pointer to the position 12 (where 'example.com' starts)
	// 	0x00, 0x01, // QTYPE (A)
	// 	0x00, 0x01, // QCLASS (IN)
	// }

	// question := []byte{
	// 	// Labels for the subdomain "sub"
	// 	0x03, 's', 'u', 'b',
	//
	// 	// Labels for the domain "example.com"
	// 	0x07, 'e', 'x', 'a', 'm', 'p', 'l', 'e',
	// 	0x03, 'c', 'o', 'm',
	//
	// 	// Compression pointer to the position after "example.com"
	// 	0xC0, 0x0B, // Assuming "example.com" starts at position 11 (0-indexed)
	//
	// 	0x00, 0x01, // QTYPE (A)
	// 	0x00, 0x01, // QCLASS (IN)
	// }

	question := []byte{156, 123, 1, 0, 0, 2, 0, 0, 0, 0, 0, 0, 3, 97, 98, 99, 17, 108, 111, 110, 103, 97, 115, 115, 100, 111, 109, 97, 105, 110, 110, 97, 109, 101, 3, 99, 111, 109, 0, 0, 1, 0, 1, 3, 100, 101, 102, 192, 16, 0, 1, 0, 1}

	labels, offset := parseDomainName(question[12:], 0)

	log.Println(offset)
	log.Println(labels)
}
