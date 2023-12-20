package main

import "encoding/binary"

type Answer struct {
	Name  string
	Type  int
	Class int
	TTL   int32
	Data  string // this is contain the sender ip address

	length int // will be the length of the data
}

func int32ToBytes(i int32) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(i))

	return res
}

func (a *Answer) Bytes() []byte {
	res := make([]byte, 0)

	name := labelSequence(a.Name)
	typebytes := intToBytes(a.Type)
	classbytes := intToBytes(a.Class)
	ttl := int32ToBytes(a.TTL)
	length := intToBytes(len(a.Data))
	data := labelSequence(a.Data)

	res = append(res, name...)
	res = append(res, typebytes...)
	res = append(res, classbytes...)
	res = append(res, ttl...)
	res = append(res, length...)
	res = append(res, data...)

	return res
}
