package playwright

type artifactImpl struct {
	channelOwner
}

func (a *artifactImpl) AbsolutePath() string { _ = "STUB: not implemented"; return "" }

func (a *artifactImpl) PathAfterFinished() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (a *artifactImpl) SaveAs(path string) error { _ = "STUB: not implemented"; return nil }

func (a *artifactImpl) Failure() error { _ = "STUB: not implemented"; return nil }

func (a *artifactImpl) Delete() error { _ = "STUB: not implemented"; return nil }

func (a *artifactImpl) Cancel() error { _ = "STUB: not implemented"; return nil }

func (a *artifactImpl) ReadIntoBuffer() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func newArtifact(parent *channelOwner, objectType string, guid string, initializer map[string]any) *artifactImpl {
	_ = "STUB: not implemented"
	return nil
}
