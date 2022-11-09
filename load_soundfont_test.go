package tsf

import (
	"fmt"
	"testing"
)

func TestLoadSoundfont(t *testing.T) {
	sf := NewSoundFontFromPath("test/Drama_Piano.sf2")
	if sf == nil {
		t.Error("soundfont is nil!")
	}
	defer sf.Close()

	fmt.Println(sf.Presets(), sf.PresetName(0))
}
