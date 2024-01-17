//go:build !nowebgl2backend && js

package auto

import "github.com/BenLubar/theseus/backend/webgl2"

var _ = addGraphicsCandidate("webgl2", webgl2.Graphics, 120)
