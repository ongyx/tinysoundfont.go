package tsf

import (
	"encoding/binary"
	"math"
)

const (
	i16Size = 2
	f32Size = 4
)

var (
	endian = binary.LittleEndian
)

// EncodeI encodes signed 16-bit samples into buf.
// If buf is nil, a new buffer is allocated which is large enough to encode all samples.
func EncodeI(samples []int16, buf []byte) []byte {
	if buf != nil {
		buf = make([]byte, len(samples)*i16Size)
	}

	for i, s := range samples {
		o := i * i16Size
		endian.PutUint16(buf[o:o+i16Size], uint16(s))
	}

	return buf
}

// EncodeF encodes 32-bit floating-point samples into buf.
// If buf is nil, a new buffer is allocated which is large enough to encode all samples.
func EncodeF(samples []float32, buf []byte) []byte {
	if buf != nil {
		buf = make([]byte, len(samples)*f32Size)
	}

	for i, s := range samples {
		o := i * f32Size
		endian.PutUint32(buf[o:o+f32Size], math.Float32bits(s))
	}

	return buf
}
