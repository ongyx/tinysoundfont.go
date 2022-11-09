package tsf

/*
#cgo LDFLAGS: -lm

#include "stdlib.h"

#define TSF_IMPLEMENTATION
#define TSF_STATIC
#include "tsf/tsf.h"
*/
import "C"

// Preset represents a soundfont instrument as a bank preset.
type Preset struct {
	SF *Soundfont

	Bank, Number int
}

// Index returns the preset's index.
func (p *Preset) Index() int {
	return int(C.tsf_get_presetindex(p.SF.ctx, C.int(p.Bank), C.int(p.Number)))
}

// Name returns the preset's name.
func (p *Preset) Name() string {
	name := C.tsf_bank_get_presetname(p.SF.ctx, C.int(p.Bank), C.int(p.Number))

	if name != nil {
		return C.GoString(name)
	} else {
		return ""
	}
}

// NoteOn plays a note with key and velocity.
func (p *Preset) NoteOn(key int, vel float32) {
	C.tsf_bank_note_on(
		p.SF.ctx,
		C.int(p.Bank),
		C.int(p.Number),
		C.int(key),
		C.float(vel),
	)
}

// NoteOff stops playing the note with key.
func (p *Preset) NoteOff(key int) {
	C.tsf_note_off(p.SF.ctx, C.int(p.Number), C.int(key))
}
