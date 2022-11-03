package tsf

/*
#cgo LDFLAGS: -lm

#include "stdlib.h"

#define TSF_IMPLEMENTATION
#define TSF_STATIC
#include "vendor/tsf/tsf.h"
*/
import "C"

const (
	// StereoInterleaved renders by alternating stereo channel samples.
	StereoInterleaved OutputMode = C.TSF_STEREO_INTERLEAVED

	// StereoUnweaved renders the samples for the left channel,
	// followed by the right.
	StereoUnweaved OutputMode = C.TSF_STEREO_INTERLEAVED

	// Mono renders by mixing the samples into a single channel.
	Mono OutputMode = C.TSF_MONO
)

// OutputMode specifies how to render samples.
type OutputMode int
