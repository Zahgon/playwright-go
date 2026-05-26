package playwright

type writableStream struct {
	channelOwner
}

func (s *writableStream) Copy(file string) error { _ = "STUB: not implemented"; return nil }

func newWritableStream(parent *channelOwner, objectType string, guid string, initializer map[string]any) *writableStream {
	_ = "STUB: not implemented"
	return nil
}
