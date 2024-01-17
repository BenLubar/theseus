package null

import (
	"flag"
)

var Audio nullAudio

type nullAudio struct{}

var flagNoNullAudio = flag.Bool("nonullaudio", false, "instead of using the null audio backend, crash")

func (nullAudio) Supported() bool {
	return !*flagNoNullAudio
}

func (nullAudio) Init() error {
	return nil
}

func (nullAudio) Shutdown() error {
	return nil
}
