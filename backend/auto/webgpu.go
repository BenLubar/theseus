//go:build !nowebgpubackend && js

package auto

import "github.com/BenLubar/theseus/backend/webgpu"

var _ = addGraphicsCandidate("webgpu", webgpu.Graphics, 200)
