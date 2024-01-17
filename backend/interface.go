package backend

// GameInterface is a set of methods that return relevant interfaces that are
// implemented by a game (as opposed to by the engine).
type GameInterface interface {
	Init() error
	Shutdown() error
}

// AudioInterface is a set of methods that must be implemented by an audio
// backend.
type AudioInterface interface {
	Supported() bool
	Init() error
	Shutdown() error
}

// GraphicsInterface is a set of methods that must be implemented by a graphics
// backend.
type GraphicsInterface interface {
	Supported() bool
	Init() error
	Shutdown() error
}
