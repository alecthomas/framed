package framed

import (
	"encoding/binary"
	"io"
)

// A Framed transport over an underlying stream.
type Framed struct {
	rw io.ReadWriter
}

// NewFramed creates a new framed transport over a stream.
func NewFramed(rw io.ReadWriter) *Framed {
	return &Framed{rw: rw}
}

// Read a single frame.
func (f *Framed) Read() ([]byte, error) {
	var size uint32
	if err := binary.Read(f.rw, binary.BigEndian, &size); err != nil {
		return nil, err
	}
	buf := make([]byte, size)
	if _, err := io.ReadFull(f.rw, buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// Write a single frame.
func (f *Framed) Write(b []byte) error {
	size := uint32(len(b))
	if err := binary.Write(f.rw, binary.BigEndian, size); err != nil {
		return err
	}
	_, err := f.rw.Write(b)
	return err
}

// Close the underlying stream (if closable).
func (f *Framed) Close() error {
	return f.rw.(io.Closer).Close()
}
