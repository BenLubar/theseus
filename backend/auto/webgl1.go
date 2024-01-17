//go:build !nowebgl1backend && js

package auto

import "github.com/BenLubar/theseus/backend/webgl1"

var _ = addGraphicsCandidate("webgl1", webgl1.Graphics, 110)
