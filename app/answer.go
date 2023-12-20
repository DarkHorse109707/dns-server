package main

type Answer struct {
	Name  string
	Type  int
	Class int
	TTL   int32
	Data  string // this is contain the sender ip address

	Length int // will be the length of the data
}

func (a *Answer) Bytes() []byte {
	res := make([]byte, 0)

	name := labelSequence(a.Name)
	typebytes := intToBytes(a.Type)
	classbytes := intToBytes(a.Class)
	ttl := int32ToBytes(a.TTL)
	length := intToBytes(a.Length)
	data := ParseIP(a.Data)

	res = append(res, name...)
	res = append(res, typebytes...)
	res = append(res, classbytes...)
	res = append(res, ttl...)
	res = append(res, length...)
	res = append(res, data...)

	return res
}
