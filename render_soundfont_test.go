package tsf

import (
	"encoding/binary"
	"testing"
)

const sampleRate = 44100

func TestRenderSoundFont(t *testing.T) {
	sf := NewSoundFontFromPath("test/Drama_Piano.sf2")
	if sf == nil {
		t.Error("soundfont is nil!")
	}
	defer sf.Close()

	sf.SetOutput(Mono, sampleRate, 0)

	sf.NoteOn(0, 60, 1)

	samples := make([]int16, sampleRate)
	sf.RenderInt(samples, false)

	ab := NewAudioBuffer()
	ab.EncodeI16(samples, binary.LittleEndian)
}
