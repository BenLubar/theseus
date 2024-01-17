package example

// Interface is a singleton implemention of backend.GameInterface.
var Interface impl

type impl struct{}

func (impl) Init() error {
	return nil
}

func (impl) Shutdown() error {
	return nil
}
