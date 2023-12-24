package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func int32ToBytes(i int32) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(i))

	return res
}

func ParseIP(ip string) []byte {
	var response []byte
	elements := strings.Split(ip, ".")

	for _, element := range elements {
		num, err := strconv.Atoi(element)
		if err != nil {
			fmt.Println("Error converting string to int: ", err)
		}
		response = append(response, byte(num))
	}
	return response
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

// this function return the domain name and the offset of the next byte
func parseDomainName(data []byte, offset int) (string, int) {
	var res strings.Builder

	log.Println("data: ", data, "data length: ", len(data), "offset: ", offset)

	// WARN: so far we know we encounter problem with wrong pointer.
	// the wrong "pointer" here's mean it's pointing into the wrong index.
	// pointer offset extracting seems being the cause.
	for {
		length := int(data[offset])
		log.Println("length: ", length)
		offset++

		if offset >= len(data) {
			break
		}

		if length == 0 {
			break // end of domain
		}

		if length >= 192 { // Pointer encountered
			log.Println("offset: ", offset)
			pointerOffset := ((int(data[offset-1]) & 0x3F) << 8) | int(data[offset])
			log.Println("pointerOffset: ", pointerOffset)

			subdomain, _ := parseDomainName(data, pointerOffset)
			log.Println("subdomain: ", subdomain)
			res.WriteString(subdomain)
			offset += 2
			break
		}

		if res.Len() > 0 {
			res.WriteByte('.')
		}

		label := data[offset : offset+length]
		res.Write(label)
		offset += length
	}

	return res.String(), offset
}

func intToBytes(n int) []byte {
	val := make([]byte, 2)

	binary.BigEndian.PutUint16(val, uint16(n))

	return val
}

func toBytes(s string) []byte {
	return []byte(s)
}
