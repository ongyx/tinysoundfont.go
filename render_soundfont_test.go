package tsf

import (
	"encoding/binary"
	"testing"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const sampleRate = 44100

func TestRenderSoundFont(t *testing.T) {
	sf := NewSoundFontFromPath("test/Drama_Piano.sf2")
	if sf == nil {
		t.Error("soundfont is nil!")
	}
	defer sf.Close()

	sf.SetOutput(StereoInterleaved, sampleRate, 0)

	sf.NoteOn(0, 60, 1)

	samples := make([]int16, sampleRate)
	sf.RenderInt(samples, false)

	var b Buffer
	b.EncodeInt(samples, binary.LittleEndian)

	c := audio.NewContext(sampleRate)

	// NOTE: we have to busy-wait so the process doesn't exit before the context initializes and the player finishes playing.
	for {
		if c.IsReady() {
			break
		}
		time.Sleep(time.Second)
	}

	p := c.NewPlayerFromBytes(b.Slice)
	p.Play()

	for {
		if !p.IsPlaying() {
			break
		}
		time.Sleep(time.Second)
	}
}
