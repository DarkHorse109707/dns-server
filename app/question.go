package main

import (
	"encoding/binary"
	"log"
)

type Question struct {
	Name  string
	Type  int
	Class int
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

func (q *Question) Parse(data []byte) {
	var offset int

	// log.Println("data: ", string(data[12:]))
	q.Name, offset = parseDomainName(data[12:], 0)
	log.Println("name: ", q.Name, "offset: ", offset)

	q.Type = int(binary.BigEndian.Uint16(data[offset : offset+2]))
	q.Class = int(binary.BigEndian.Uint16(data[offset+2 : offset+4]))

	log.Printf("QUESTION: %+v\n", q)
}
