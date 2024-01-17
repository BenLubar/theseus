//go:build js

package webgl2

import "syscall/js"

var Graphics graphics

type graphics struct{}

func (graphics) Supported() bool {
	if !js.Global().Get("WebGL2RenderingContext").Truthy() {
		return false
	}

	canvas := js.Global().Get("document").Call("createElement", "canvas")
	ctx := canvas.Call("getContext", "webgl2")
	return ctx.Truthy()
}

func (graphics) Init() error {
	return nil
}

func (graphics) Shutdown() error {
	return nil
}
