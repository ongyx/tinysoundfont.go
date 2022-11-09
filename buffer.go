package tsf

import (
	"bytes"
	"encoding/binary"
	"math"
)

const (
	f32Size = 4
	i16Size = 2
)

// AudioBuffer is a buffer for encoding samples as bytes.
//
// NOTE: All encode methods reset the buffer before encoding the samples.
type AudioBuffer struct {
	*bytes.Buffer
}

// F32Encode encodes the float32 samples into the buffer.
func (ab AudioBuffer) F32Encode(samples []float32, ord binary.ByteOrder) {
	ab.Reset()
	ab.Grow(len(samples) * f32Size)

	temp := make([]byte, f32Size)

	for _, s := range samples {
		ord.PutUint32(temp, math.Float32bits(s))
		ab.Write(temp)
	}
}

// I16Encode encodes the int16 samples into the buffer.
func (ab AudioBuffer) I16Encode(samples []int16, ord binary.ByteOrder) {
	ab.Reset()
	ab.Grow(len(samples) * i16Size)

	temp := make([]byte, i16Size)

	for _, s := range samples {
		ord.PutUint16(temp, uint16(s))
		ab.Write(temp)
	}
}
