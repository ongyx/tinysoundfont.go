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

// NewAudioBuffer creates an empty audio buffer.
func NewAudioBuffer() AudioBuffer {
	return AudioBuffer{new(bytes.Buffer)}
}

// EncodeF32 encodes the float32 samples into the buffer.
func (ab AudioBuffer) EncodeF32(samples []float32, ord binary.ByteOrder) {
	ab.Reset()
	ab.Grow(len(samples) * f32Size)

	temp := make([]byte, f32Size)

	for _, s := range samples {
		ord.PutUint32(temp, math.Float32bits(s))
		ab.Write(temp)
	}
}

// EncodeI16 encodes the int16 samples into the buffer.
func (ab AudioBuffer) EncodeI16(samples []int16, ord binary.ByteOrder) {
	ab.Reset()
	ab.Grow(len(samples) * i16Size)

	temp := make([]byte, i16Size)

	for _, s := range samples {
		ord.PutUint16(temp, uint16(s))
		ab.Write(temp)
	}
}
