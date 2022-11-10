# tinysoundfont.go

Go bindings for [tinysoundfont], a single header SoundFont2 synthesizer.

```go
const sampleRate = 44100

// Load soundfont from a file path
sf := tsf.NewSoundfontFromPath("soundfont.sf2")
if sf == nil {
  panic("failed to load soundfont!")
}
defer sf.Close()

// Set output to mono (single channel) and 44.1kHz.
sf.SetOutput(tsf.Mono, sampleRate, 0)

// Play MIDI note 60 (middle C) with preset 0 and velocity 1.
sf.NoteOn(0, 60, 1)

// Allocate an output buffer for half a second of the note playing.
// If SetOutput() was called with StereoInterleaved or StereoUnweaved,
// the buffer should be twice as long as the duration.
out := make([]int16, sampleRate / 2)
sf.RenderInt(out, false)
```

## Building

This package requires CGo to be enabled, which is the default for native builds.
However, you have to configure the specific compiler for your target if you are cross-compiling:

```
$ export CGO_ENABLED=1
$ export CC=(cross compiler)
$ GOOS=(target OS) GOARCH=(target arch) go build ...
```

## License

tinysoundfont.go is licensed under the MIT License.

[tinysoundfont]: https://github.com/schellingb/TinySoundFont
