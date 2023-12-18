package main

import (
	"encoding/binary"
	"strings"
)

type Question struct {
	Name  string
	Type  int
	Class int
}

func TypeNameToValue(s string) byte {
	switch s {
	case "A":
		return 1
	case "NS":
		return 2
	case "MD":
		return 3
	case "MF":
		return 4
	case "CNAME":
		return 5
	case "SOA":
		return 6
	case "MB":
		return 7
	case "MG":
		return 8
	case "MR":
		return 9
	case "NULL":
		return 10
	case "WKS":
		return 11
	case "PTR":
		return 12
	case "HINFO":
		return 13
	case "MINFO":
		return 14
	case "MX":
		return 15
	case "TXT":
		return 16
	default:
		return 0
	}
}

func ClassNameToValue(s string) byte {
	switch s {
	case "IN":
		return 1
	case "CS":
		return 2
	case "CH":
		return 3
	case "HS":
		return 4
	default:
		return 0
	}
}

func labelSequence(domain string) []byte {
	labels := strings.Split(domain, ".")
	var sequence []byte
	for _, label := range labels {
		sequence = append(sequence, byte(len(label)))
		sequence = append(sequence, label...)
	}
	sequence = append(sequence, '\x00')
	return sequence
}

func intToBytes(n int) []byte {
	val := make([]byte, 2)

	binary.BigEndian.PutUint16(val, uint16(n))

	return val
}

func (q *Question) ToBytes() []byte {
	var res []byte

	name := labelSequence(q.Name)
	typebyte := intToBytes(q.Type)
	classbyte := intToBytes(q.Class)

	res = append(res, name...)
	res = append(res, typebyte...)
	res = append(res, classbyte...)

	return res
}
