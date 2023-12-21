package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeader(t *testing.T) {
	header := &Header{
		ID:    1234,
		QR:    true,
		RCODE: 4,
	}

	newheader := new(Header)

	newheader.Parse(header.ToBytes())

	assert.Equal(t, header, newheader)
}
