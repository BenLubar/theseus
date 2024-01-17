//go:build !nowebgpubackend && js

package auto

import "github.com/BenLubar/toy-renderer-implementation/backend/webgpu"

var _ = addGraphicsCandidate("webgpu", webgpu.Graphics, 200)
