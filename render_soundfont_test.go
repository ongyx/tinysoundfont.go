package tsf

import (
	"testing"
)

func TestRenderSoundFont(t *testing.T) {
	sf := NewSoundFont("test/Drama_Piano.sf2")
	if sf == nil {
		t.Error("soundfont is nil!")
	}
	defer sf.Close()

}
