package main

import "encoding/binary"

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
