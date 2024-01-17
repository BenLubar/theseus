package auto

import "flag"

func Main() {
	flag.Parse()

	for {
		select {}
	}
}
