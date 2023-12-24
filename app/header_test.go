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
	// Combine header and question
	question := []byte{
		// DNS Header (truncated for brevity)
		0xAB, 0xCD, // ID (16 bits)
		0x01, 0x00, // Flags (QR = 0, Opcode = 0, AA = 0, TC = 0, RD = 1, RA = 0, Z = 0, RCODE = 0)
		0x00, 0x01, // QDCount (1 question)
		0x00, 0x00, // ANCount (0 answers)
		0x00, 0x00, // NSCount (0 authority records)
		0x00, 0x00, // ARCount (0 additional records)

		// DNS Question
		0x07, 'e', 'x', 'a', 'm', 'p', 'l', 'e',
		0x03, 'c', 'o', 'm',
		0xC8, 0x0C, // Compression pointer to the position where 'example.com' starts

		0x00, 0x01, // QTYPE (A)
		0x00, 0x01, // QCLASS (IN)
	}
	// question = question[12:]
	// log.Println("this is bitwise operation of 222: ", 0xDE&0x3F)

	quest := new(Question)
	quest.Parse(question)
	labels, offset := parseDomainName(question[12:], 0)
	log.Println("labels: ", labels, "offset: ", offset)

	log.Printf("quest: %+v", quest)
}
