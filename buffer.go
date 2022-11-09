package tsf

import (
	"encoding/binary"
	"math"
)

const (
	f32Size = 4
	i16Size = 2
)

// Buffer is a wrapper around a byte slice for encoding audio samples.
// The zero value of a buffer can be used, as the slice is allocated by the encode methods if nil.
type Buffer struct {
	Slice []byte
}

// EncodeInt encodes the 16-bit signed samples with the byte order into the buffer.
func (b *Buffer) EncodeInt(samples []int16, ord binary.ByteOrder) {
	b.alloc(len(samples) * i16Size)

	for i, s := range samples {
		o := i * i16Size
		ord.PutUint16(b.Slice[o:o+i16Size], uint16(s))
	}
}

// EncodeFloat encodes the 32-bit floating point samples with the byte order into the buffer.
func (b *Buffer) EncodeFloat(samples []float32, ord binary.ByteOrder) {
	b.alloc(len(samples) * f32Size)

	for i, s := range samples {
		o := i * f32Size
		ord.PutUint32(b.Slice[o:o+f32Size], math.Float32bits(s))
	}
}

func (b *Buffer) alloc(size int) {
	if b.Slice == nil {
		// allocate new buffer
		b.Slice = make([]byte, size)
	} else {
		l := len(b.Slice)
		if l < size {
			// expand buffer if not big enough to encode all samples
			b.Slice = append(b.Slice, make([]byte, size-l)...)
		} else {
			// buffer is bigger than requested size, reslice
			b.Slice = b.Slice[:size]
		}
	}
}
