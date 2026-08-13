package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

const (
	DatagramSize = 1516             // the maximum supported datagram size
	BlockSize    = DatagramSize - 4 // the DatagramSize minus a 4-byte header
)

type OpCode uint16

const (
	OpRRQ OpCode = iota + 1
	_            // no WRQ support
	OpData
	OpAck
	OpErr
)

type ErrCode uint16

const (
	ErrUnknown ErrCode = iota
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID
	ErrFileExists
	ErrNoUser
)

type Data struct {
	Block   uint16
	Payload io.Reader
}

type Ack uint16

func (a Ack) MarshalBinary() ([]byte, error) {
	cap := 2 + 2 // operation code + block number

	b := new(bytes.Buffer)
	b.Grow(cap)

	// write operation code.
	if err := binary.Write(b, binary.BigEndian, OpAck); err != nil {
		return nil, err
	}

	// write block number
	if err := binary.Write(b, binary.BigEndian, a); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (a *Ack) UnmarshalBinary(p []byte) error {
	var code OpCode

	r := bytes.NewReader(p)

	// read operation code
	if err := binary.Read(r, binary.BigEndian, &code); err != nil {
		return err
	}

	if code != OpAck {
		return errors.New("invalid ACK")
	}

	// read block number
	return binary.Read(r, binary.BigEndian, a)
}
