package main

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
