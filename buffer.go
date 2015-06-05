package minimembuf

import "io"

// Buffer implements a 2-word in-memory, garbage-littering io.ReadWriter
type Buffer string

// Read extracts byte data out of the buffer.
func (b *Buffer) Read(p []byte) (n int, err error) {
	n = copy(p, *b)
	*b = (*b)[n:]
	if len(*b) == 0 {
		return n, io.EOF
	}
	return n, nil
}

// Write stores byte data into the buffer.
func (b *Buffer) Write(p []byte) (n int, err error) {
	*b = *b + Buffer(p)
	return len(p), nil
}

// Len returns the size of available data for reading.
func (b *Buffer) Len() int { return len(*b) }

// Cap returns the total size of the buffer in memory.
func (b *Buffer) Cap() int { return b.Len() }
