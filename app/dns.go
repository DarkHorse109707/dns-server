package main

type DNS struct {
	Question *Question
	Header   *Header
	Answer   *Answer
}

func (d *DNS) Bytes() []byte {
	var res []byte

	res = append(res, d.Header.ToBytes()...)
	res = append(res, d.Question.ToBytes()...)

	return res
}
