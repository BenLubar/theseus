package auto

import (
	"flag"

	"github.com/BenLubar/toy-renderer-implementation/backend"
)

func Main(iface backend.GameInterface) {
	flag.Parse()

	for {
		select {}
	}
}
