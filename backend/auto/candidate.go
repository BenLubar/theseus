package auto

import (
	"log"
	"sort"

	"github.com/BenLubar/toy-renderer-implementation/backend"
)

type graphicsCandidate struct {
	name     string
	iface    backend.GraphicsInterface
	priority int
}

var graphicsCandidates []graphicsCandidate

func addGraphicsCandidate(name string, iface backend.GraphicsInterface, priority int) struct{} {
	graphicsCandidates = append(graphicsCandidates, graphicsCandidate{
		name:     name,
		iface:    iface,
		priority: priority,
	})

	return struct{}{}
}

type audioCandidate struct {
	name     string
	iface    backend.AudioInterface
	priority int
}

var audioCandidates []audioCandidate

func addAudioCandidate(name string, iface backend.AudioInterface, priority int) struct{} {
	audioCandidates = append(audioCandidates, audioCandidate{
		name:     name,
		iface:    iface,
		priority: priority,
	})

	return struct{}{}
}

func chooseBackends() {
	sort.Slice(graphicsCandidates, func(i, j int) bool {
		return graphicsCandidates[i].priority > graphicsCandidates[j].priority
	})
	sort.Slice(audioCandidates, func(i, j int) bool {
		return audioCandidates[i].priority > audioCandidates[j].priority
	})

	for _, c := range graphicsCandidates {
		if c.iface.Supported() {
			log.Println("using graphics backend:", c.name)
			Graphics = c.iface
			break
		}
	}
	for _, c := range audioCandidates {
		if c.iface.Supported() {
			log.Println("using audio backend:", c.name)
			Audio = c.iface
			break
		}
	}

	// let the candidate slices get garbage collected
	graphicsCandidates = nil
	audioCandidates = nil
}
