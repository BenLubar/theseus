//go:build js

package internal

import (
	"context"
	"syscall/js"
)

func Await(ctx context.Context, promise js.Value) (v js.Value, err error) {
	resolve := make(chan js.Value, 1)
	reject := make(chan js.Error, 1)

	var onResolve, onReject js.Func
	cleanup := func() {
		onResolve.Release()
		onReject.Release()
	}
	onResolve = js.FuncOf(func(this js.Value, args []js.Value) any {
		cleanup()

		resolve <- args[0]

		return js.Undefined()
	})
	onReject = js.FuncOf(func(this js.Value, args []js.Value) any {
		cleanup()

		reject <- js.Error{Value: args[0]}

		return js.Undefined()
	})

	promise.Call("then", onResolve, onReject)

	select {
	case <-ctx.Done():
		err = ctx.Err()
	case v = <-resolve:
	case err = <-reject:
	}

	return
}
