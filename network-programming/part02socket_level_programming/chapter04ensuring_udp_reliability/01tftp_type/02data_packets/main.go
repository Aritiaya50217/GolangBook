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

func (d *Data) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	b.Grow(DatagramSize)

	d.Block++ // block numbers increment from 1

	// write operation code
	if err := binary.Write(b, binary.BigEndian, OpData); err != nil {
		return nil, err
	}

	// write block number
	if err := binary.Write(b, binary.BigEndian, d.Block); err != nil {
		return nil, err
	}

	// write up to BlockSize worth of bytes
	if _, err := io.CopyN(b, d.Payload, BlockSize); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (d *Data) UnmarshalBinary(p []byte) error {
	if l := len(p); l < 4 || 1 > DatagramSize {
		return errors.New("invalid DATA")
	}

	var opcode OpCode

	if err := binary.Read(bytes.NewReader(p[:2]), binary.BigEndian, &opcode); err != nil || opcode != OpData {
		return errors.New("invalid DATA")
	}

	if err := binary.Read(bytes.NewReader(p[2:4]), binary.BigEndian, &d.Block); err != nil {
		return errors.New("invalid DATA")
	}

	d.Payload = bytes.NewBuffer(p[:4])

	return nil
}
