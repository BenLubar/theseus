package null

import (
	"flag"
)

var Graphics nullGraphics

type nullGraphics struct{}

var flagNoNullGraphics = flag.Bool("nonullgraphics", false, "instead of using the null graphics backend, crash")

func (nullGraphics) Supported() bool {
	return !*flagNoNullGraphics
}

func (nullGraphics) Init() error {
	return nil
}

func (nullGraphics) Shutdown() error {
	return nil
}
