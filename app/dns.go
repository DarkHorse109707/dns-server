package main

type DNS struct {
	Question []*Question
	Header   *Header
	Answer   []*Answer
}

func (d *DNS) Bytes() []byte {
	var res []byte

	res = append(res, d.Header.ToBytes()...)

	for _, v := range d.Question {
		res = append(res, v.ToBytes()...)
	}

	for _, v := range d.Answer {
		res = append(res, v.Bytes()...)
	}

	return res
}

func NewDNS(data []byte) *DNS {
	header := new(Header)
	header.Parse(data[:12])

	questions := make([]*Question, 0, header.QDCOUNT)
	answers := make([]*Answer, 0, header.QDCOUNT)
	n := 12
	for i := 0; i < int(header.QDCOUNT); i++ {
		var q *Question
		q, n = NewQuestion(data, n)
		questions = append(questions, q)
		answers = append(answers, &Answer{
			Name:   q.Name,
			Type:   q.Type,
			Class:  q.Class,
			TTL:    60,
			Length: 4,
			Data:   "8.8.8.8",
		})
	}

	header.ANCOUNT = uint16(len(answers))
	return &DNS{
		Question: questions,
		Header:   header,
		Answer:   answers,
	}
}
