package playwright

type streamImpl struct {
	channelOwner
}

func (s *streamImpl) SaveAs(path string) error { _ = "STUB: not implemented"; return nil }

func (s *streamImpl) ReadAll() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func newStream(parent *channelOwner, objectType string, guid string, initializer map[string]any) *streamImpl {
	_ = "STUB: not implemented"
	return nil
}
