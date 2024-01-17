//go:build js

package webgpu

import (
	"context"
	"syscall/js"

	"github.com/BenLubar/toy-renderer-implementation/internal"
)

var Graphics graphics

type graphics struct{}

func (graphics) Supported() bool {
	gpu := js.Global().Get("navigator").Get("gpu")
	if !gpu.Truthy() {
		return false
	}

	adapter, _ := internal.Await(context.Background(), gpu.Call("requestAdapter"))

	return adapter.Truthy()
}

func (graphics) Init() error {
	return nil
}

func (graphics) Shutdown() error {
	return nil
}
