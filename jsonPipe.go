package playwright

type jsonPipe struct {
	channelOwner
	msgChan chan *message
}

func (j *jsonPipe) Send(message map[string]any) error { _ = "STUB: not implemented"; return nil }

func (j *jsonPipe) Close() error { _ = "STUB: not implemented"; return nil }

func (j *jsonPipe) Poll() (*message, error) { _ = "STUB: not implemented"; return nil, nil }

func newJsonPipe(parent *channelOwner, objectType string, guid string, initializer map[string]any) *jsonPipe {
	_ = "STUB: not implemented"
	return nil
}

// Send directly to maintain message ordering - the channel buffer prevents blocking
// Previously used a goroutine which could cause out-of-order delivery

// Recover from panic if channel is closed
