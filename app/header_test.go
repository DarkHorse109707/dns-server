package main

import (
	"fmt"
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
