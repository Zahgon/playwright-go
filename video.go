package playwright

import (
	"sync"
)

type videoImpl struct {
	page         *pageImpl
	artifact     *artifactImpl
	artifactChan chan *artifactImpl
	done         chan struct{}
	closeOnce    sync.Once
	isRemote     bool
}

func (v *videoImpl) Path() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (v *videoImpl) Delete() error { _ = "STUB: not implemented"; return nil }

func (v *videoImpl) SaveAs(path string) error { _ = "STUB: not implemented"; return nil }

func (v *videoImpl) artifactReady(artifact *artifactImpl) { _ = "STUB: not implemented"; return }

func (v *videoImpl) pageClosed(p Page) { _ = "STUB: not implemented"; return }

func (v *videoImpl) getArtifact() {
	_ = "STUB: not implemented"
	// prevent channel block if no video will be produced
	return
}

// no recordVideo option

// page closed
// make sure get artifact if it's ready before page closed

func newVideo(page *pageImpl) *videoImpl { _ = "STUB: not implemented"; return nil }
