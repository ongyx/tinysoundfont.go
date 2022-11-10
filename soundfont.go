package tsf

/*
#cgo LDFLAGS: -lm

#include "stdlib.h"

#define TSF_IMPLEMENTATION
#define TSF_STATIC
#include "tsf/tsf.h"
*/
import "C"

import "unsafe"

// Soundfont represents a SF2 file in memory.
type Soundfont struct {
	ctx  *C.tsf
	mono bool
}

// NewSoundFont loads a soundfont from a buffer.
// If the soundfont fails to load, nil is returned.
func NewSoundFont(buf []byte) *Soundfont {
	cbuf := unsafe.Pointer(&buf[0])

	if p := C.tsf_load_memory(cbuf, C.int(len(buf))); p != nil {
		return &Soundfont{ctx: p}
	} else {
		return nil
	}
}

// NewSoundFontFromPath loads a soundfont from the given path.
// If the soundfont fails to load, nil is returned.
func NewSoundFontFromPath(path string) *Soundfont {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	if p := C.tsf_load_filename(cpath); p != nil {
		return &Soundfont{ctx: p}
	} else {
		return nil
	}
}

// Copy copies the soundfont into a separate instance.
// Any copies should be closed with Close() after use.
func (sf *Soundfont) Copy() *Soundfont {
	return &Soundfont{ctx: C.tsf_copy(sf.ctx)}
}

// Close frees the memory used by the soundfont.
// The soundfont should not be used past this point.
func (sf *Soundfont) Close() {
	C.tsf_close(sf.ctx)
	sf.ctx = nil
}

// Reset stops all playing notes and resets all channel parameters.
func (sf *Soundfont) Reset() {
	C.tsf_reset(sf.ctx)
}

// Presets returns the number of presets in the soundfont.
func (sf *Soundfont) Presets() int {
	return int(C.tsf_get_presetcount(sf.ctx))
}

// PresetName returns the name of a preset index.
func (sf *Soundfont) PresetName(index int) string {
	name := C.tsf_get_presetname(sf.ctx, C.int(index))

	if name != nil {
		return C.GoString(name)
	} else {
		return ""
	}
}

// SetOutput sets the parameters for rendering voices.
// samplerate is in hertz (typically 44100) and gain is in decibels.
func (sf *Soundfont) SetOutput(mode OutputMode, samplerate int, gain float32) {
	sf.mono = mode == Mono

	C.tsf_set_output(
		sf.ctx,
		C.enum_TSFOutputMode(mode),
		C.int(samplerate),
		C.float(gain),
	)
}

// SetVolume sets the global gain factor (1.0 = 100%).
func (sf *Soundfont) SetVolume(gain float32) {
	C.tsf_set_volume(sf.ctx, C.float(gain))
}

// SetMaxVoices sets the maximum number of voices that can play simultaneously.
func (sf *Soundfont) SetMaxVoices(num int) {
	C.tsf_set_max_voices(sf.ctx, C.int(num))
}

// NoteOn plays a note given the preset index, note key and velocity.
func (sf *Soundfont) NoteOn(index, key int, velocity float32) {
	C.tsf_note_on(
		sf.ctx,
		C.int(index),
		C.int(key),
		C.float(velocity),
	)
}

// NoteOff stops playing a note.
func (sf *Soundfont) NoteOff(index, key int) {
	C.tsf_note_off(sf.ctx, C.int(index), C.int(key))
}

// NoteOffAll stops playing all notes.
func (sf *Soundfont) NoteOffAll() {
	C.tsf_note_off_all(sf.ctx)
}

// ActiveVoices returns the number of voices allocated
// (i.e how many notes can play simultaneously).
func (sf *Soundfont) ActiveVoices() int {
	return int(C.tsf_active_voice_count(sf.ctx))
}

// RenderInt renders the notes as 16-bit signed samples into a buffer.
// If mix is true, the samples are mixed into the buffer's existing samples.
func (sf *Soundfont) RenderInt(buf []int16, mix bool) {
	var mixFlag int
	if mix {
		mixFlag = 1
	}

	size := len(buf)
	if !sf.mono {
		size /= 2
	}

	C.tsf_render_short(
		sf.ctx,
		(*C.short)(&buf[0]),
		C.int(size),
		C.int(mixFlag),
	)
}

// RenderFloat renders the notes as 32-bit floating-point samples into a buffer.
// If mix is true, the samples are mixed into the buffer's existing samples.
func (sf *Soundfont) RenderFloat(buf []float32, mix bool) {
	var mixFlag int
	if mix {
		mixFlag = 1
	}

	size := len(buf)
	if !sf.mono {
		size /= 2
	}

	C.tsf_render_float(
		sf.ctx,
		(*C.float)(&buf[0]),
		C.int(size),
		C.int(mixFlag),
	)
}
