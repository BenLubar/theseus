//go:build !nonullbackend

package auto

import "github.com/BenLubar/toy-renderer-implementation/backend/null"

var (
	_ = addGraphicsCandidate("null", null.Graphics, 0)
	_ = addAudioCandidate("null", null.Audio, 0)
)
