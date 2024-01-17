//go:build js

package webgl1

import "syscall/js"

var Graphics graphics

type graphics struct{}

func (graphics) Supported() bool {
	if !js.Global().Get("WebGLRenderingContext").Truthy() {
		return false
	}

	canvas := js.Global().Get("document").Call("createElement", "canvas")
	ctx := canvas.Call("getContext", "webgl")
	return ctx.Truthy()
}

func (graphics) Init() error {
	return nil
}

func (graphics) Shutdown() error {
	return nil
}
