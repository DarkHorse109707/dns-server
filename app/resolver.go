package main

import (
	"net"
	"time"
)

const WRITE_READ_DEADLINE = 15

type Resolver struct {
	Conn *net.UDPConn
	*net.UDPAddr
}

func NewResolver(resolver string) (*Resolver, error) {
	addr, err := net.ResolveUDPAddr("udp", resolver)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return &Resolver{
		Conn:    conn,
		UDPAddr: addr,
	}, nil
}

// TODO: refactor to satisfied the need "for each outgoing packet to the resolver make sure it has only one question."
// this function returns a result from the resolver address.
func (r *Resolver) Serialize(dns *DNS) (*DNS, error) {
	res := new(DNS)
	res.Header = dns.Header
	res.Question = dns.Question

	for _, q := range dns.Question {
		dnsresolver := MessageSerialize(dns.Header, q)

		if _, err := r.write(dnsresolver.Bytes()); err != nil {
			return nil, err
		}

		buf := make([]byte, 512)

		n, _, err := r.read(buf)
		if err != nil {
			return nil, err
		}

		dnsresolver = NewDNS(buf[:n], "8.8.8.8")

		res.Answer = append(res.Answer, dnsresolver.Answer...)
	}

	return res, nil
}

func (r *Resolver) write(msg []byte) (int, error) {
	if err := r.Conn.SetWriteDeadline(time.Now().Add(WRITE_READ_DEADLINE * time.Second)); err != nil {
		return 0, err
	}

	return r.Conn.WriteToUDP(msg, r.UDPAddr)
}

func (r *Resolver) read(buf []byte) (int, *net.UDPAddr, error) {
	if err := r.Conn.SetReadDeadline(time.Now().Add(WRITE_READ_DEADLINE * time.Second)); err != nil {
		return 0, nil, err
	}

	return r.Conn.ReadFromUDP(buf)
}

func MessageSerialize(h *Header, q *Question) *DNS {
	return &DNS{
		Question: []*Question{q},
		Header: &Header{
			ID:      h.ID,
			QDCOUNT: 1,
			ANCOUNT: 0,
			NSCOUNT: 0,
			ARCOUNT: 0,
			QR:      false,
			OPCODE:  h.OPCODE,
			AA:      false,
			TC:      false,
			RD:      h.RD,
			RA:      false,
			Z:       0,
			RCODE:   h.RCODE,
		},
		Answer: make([]*Answer, 0),
	}
}
