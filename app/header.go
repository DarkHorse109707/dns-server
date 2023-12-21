package main

import (
	"encoding/binary"
)

type Header struct {
	ID      uint16
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16

	QR     bool
	OPCODE uint8 // this is should occupied 4 bit
	AA     bool
	TC     bool
	RD     bool
	RA     bool
	Z      uint8 // this is should occupied 3 bit
	RCODE  uint8 // this is should occupied 4 bit
}

func BoolToByte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func (h *Header) ToBytes() []byte {
	val := make([]byte, 12)

	binary.BigEndian.PutUint16(val[0:2], h.ID)
	val[2] = BoolToByte(h.QR)<<7 | h.OPCODE<<3 | BoolToByte(h.AA)<<2 | BoolToByte(h.TC)<<1 | BoolToByte(h.RD)
	val[3] = BoolToByte(h.RA)<<7 | h.Z<<4 | h.RCODE
	binary.BigEndian.PutUint16(val[4:6], h.QDCOUNT)
	binary.BigEndian.PutUint16(val[6:8], h.ANCOUNT)
	binary.BigEndian.PutUint16(val[8:10], h.NSCOUNT)
	binary.BigEndian.PutUint16(val[10:12], h.ARCOUNT)

	return val
}

func (h *Header) Parse(data []byte) {
	h.ID = binary.BigEndian.Uint16(data[0:2])

	// length 2 section
	h.QR = (data[2]>>7)&1 == 1
	h.OPCODE = (data[2] >> 3) & 0b00001111
	h.AA = (data[2]>>2)&1 == 1
	h.TC = (data[2]>>1)&1 == 1
	h.RD = data[2]&1 == 1

	// length 3 section
	h.RA = (data[3]>>7)&1 == 1
	h.Z = (data[3] >> 4) & 0b00000111
	RCODE := data[3] & 0b00001111
	if h.OPCODE == 0 {
		RCODE = 0
	} else {
		RCODE = 4
	}
	h.RCODE = RCODE

	h.QDCOUNT = binary.BigEndian.Uint16(data[4:6])
	h.ANCOUNT = binary.BigEndian.Uint16(data[6:8])
	h.NSCOUNT = binary.BigEndian.Uint16(data[8:10])
	h.ARCOUNT = binary.BigEndian.Uint16(data[10:12])
}
