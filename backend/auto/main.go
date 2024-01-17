package auto

import (
	"flag"
	"fmt"

	"github.com/BenLubar/toy-renderer-implementation/backend"
)

// Chosen interfaces. Only valid once Main has been called.
var (
	Audio    backend.AudioInterface
	Graphics backend.GraphicsInterface
)

// Main decides a set of backends, then runs the program using the provided
// backend.GameInterface. The backend variables in this package are not valid
// until Main has been called. Once Main returns, the program should exit.
func Main(Game backend.GameInterface) (err error) {
	flag.Parse()

	chooseBackends()

	defer doInitShutdown(Graphics, "graphics backend", &err)()
	defer doInitShutdown(Audio, "audio backend", &err)()
	defer doInitShutdown(Game, "game logic", &err)()
	if err != nil {
		return
	}

	for {
		select {}
	}
}

func doInit(iface interface{ Init() error }, name string, err *error) {
	if e := iface.Init(); *err == nil {
		*err = fmt.Errorf("initializing %s failed: %w", name, e)
	}
}

func doShutdown(iface interface{ Shutdown() error }, name string, err *error) {
	if e := iface.Shutdown(); *err == nil {
		*err = fmt.Errorf("shutting down %s failed: %w", name, e)
	}
}

func doInitShutdown(iface interface {
	Init() error
	Shutdown() error
}, name string, err *error) func() {
	doInit(iface, name, err)

	return func() {
		doShutdown(iface, name, err)
	}
}
