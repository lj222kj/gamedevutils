package gogamedevutils

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Datagram struct {
	Id int8
}

func (p *Datagram) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, p); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func FromBytes(b []byte) (*Datagram, error) {
	buf := bytes.NewBuffer(b)

	var p Datagram

	if err := binary.Read(buf, binary.LittleEndian, &p); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &p, nil
}
