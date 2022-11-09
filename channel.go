package tsf

/*
#cgo LDFLAGS: -lm

#include "stdlib.h"

#define TSF_IMPLEMENTATION
#define TSF_STATIC
#include "tsf/tsf.h"
*/
import "C"

// Channel is a note track.
// Any parameters set on the channel do not affect the underlying soundfont.
type Channel struct {
	SF *Soundfont

	ID int
}

// PresetIndex returns the preset index.
func (c *Channel) PresetIndex() int {
	return int(C.tsf_channel_get_preset_index(c.SF.ctx, C.int(c.ID)))
}

// SetPresetIndex sets the preset index.
func (c *Channel) SetPresetIndex(index int) {
	C.tsf_channel_set_presetindex(c.SF.ctx, C.int(c.ID), C.int(index))
}

// Preset returns the bank preset.
func (c *Channel) Preset() int {
	return int(C.tsf_channel_get_preset_number(c.SF.ctx, C.int(c.ID)))
}

// SetPreset sets the bank preset.
// If drum is true, the preset is treated as a drum instrument.
func (c *Channel) SetPreset(preset int, drum bool) {
	var drumFlag int
	if drum {
		drumFlag = 1
	}

	C.tsf_channel_set_presetnumber(c.SF.ctx, C.int(c.ID), C.int(preset), C.int(drumFlag))
}

// Bank returns the bank.
func (c *Channel) Bank() int {
	return int(C.tsf_channel_get_preset_bank(c.SF.ctx, C.int(c.ID)))
}

// SetBank sets the bank.
func (c *Channel) SetBank(bank int) {
	C.tsf_channel_set_bank(c.SF.ctx, C.int(c.ID), C.int(bank))
}

// SetBankPreset sets the bank and bank preset.
func (c *Channel) SetBankPreset(bank, preset int) {
	C.tsf_channel_set_bank_preset(c.SF.ctx, C.int(c.ID), C.int(bank), C.int(preset))
}

// Pan returns the stereo panning value
// (left = 0.0, center = 0.5 (default), right = 1.0).
func (c *Channel) Pan() float32 {
	return float32(C.tsf_channel_get_pan(c.SF.ctx, C.int(c.ID)))
}

// SetPan sets the stereo panning value.
func (c *Channel) SetPan(pan float32) {
	C.tsf_channel_set_pan(c.SF.ctx, C.int(c.ID), C.float(pan))
}

// Volume returns the volume (muted = 0.0, default = 1.0)
func (c *Channel) Volume() float32 {
	return float32(C.tsf_channel_get_volume(c.SF.ctx, C.int(c.ID)))
}

// SetVolume sets the volume.
func (c *Channel) SetVolume(volume float32) {
	C.tsf_channel_set_volume(c.SF.ctx, C.int(c.ID), C.float(volume))
}

// PitchWheel returns the pitch wheel position
// (min = 0, unpitched (default) = 8192, max = 16383).
func (c *Channel) PitchWheel() int {
	return int(C.tsf_channel_get_pitchwheel(c.SF.ctx, C.int(c.ID)))
}

// SetPitchWheel sets the pitch wheel position.
func (c *Channel) SetPitchWheel(pitchWheel int) {
	C.tsf_channel_set_pitchwheel(c.SF.ctx, C.int(c.ID), C.int(pitchWheel))
}

// PitchRange returns the range of the pitch wheel in semitones (default = 2.0).
func (c *Channel) PitchRange() float32 {
	return float32(C.tsf_channel_get_pitchrange(c.SF.ctx, C.int(c.ID)))
}

// SetPitchRange sets the range of the pitch wheel.
func (c *Channel) SetPitchRange(pitchRange float32) {
	C.tsf_channel_set_pitchrange(c.SF.ctx, C.int(c.ID), C.float(pitchRange))
}

// Tuning returns the voice tuning (default = 0.0, i.e A440).
func (c *Channel) Tuning() float32 {
	return float32(C.tsf_channel_get_tuning(c.SF.ctx, C.int(c.ID)))
}

// SetTuning sets the voice tuning.
func (c *Channel) SetTuning(tuning float32) {
	C.tsf_channel_set_tuning(c.SF.ctx, C.int(c.ID), C.float(tuning))
}

// NoteOn plays the note with key and velocity.
func (c *Channel) NoteOn(key int, vel float32) {
	C.tsf_channel_note_on(c.SF.ctx, C.int(c.ID), C.int(key), C.float(vel))
}

// NoteOff stops playing the note.
func (c *Channel) NoteOff(key int) {
	C.tsf_channel_note_off(c.SF.ctx, C.int(c.ID), C.int(key))
}

// NoteOffAll stops playing all notes.
func (c *Channel) NoteOffAll() {
	C.tsf_channel_note_off_all(c.SF.ctx, C.int(c.ID))
}

// SoundOffAll stops playing all notes without sustain and release.
func (c *Channel) SoundOffAll() {
	C.tsf_channel_sounds_off_all(c.SF.ctx, C.int(c.ID))
}

// MIDIControl applies a MIDI control value to a controller.
func (c *Channel) MIDIControl(ctl, value int) {
	C.tsf_channel_midi_control(c.SF.ctx, C.int(c.ID), C.int(ctl), C.int(value))
}
