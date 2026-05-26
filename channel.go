package playwright

type channel struct {
	eventEmitter
	guid       string
	connection *connection
	owner      *channelOwner // to avoid type conversion
	object     any           // retain type info (for fromChannel needed)
}

func (c *channel) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// for catch errors of route handlers etc.
func (c *channel) CreateTask(fn func()) { _ = "STUB: not implemented"; return }

func (c *channel) Send(method string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// GUIDs are now always eagerly resolved in connection.Dispatch

func (c *channel) SendReturnAsDict(method string, options ...any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GUIDs are now always eagerly resolved in connection.Dispatch

func (c *channel) innerSend(method string, options ...any) *protocolCallback {
	_ = "STUB: not implemented"
	return nil
}

// SendNoReply ignores return value and errors
// almost equivalent to `send(...).catch(() => {})`
func (c *channel) SendNoReply(method string, options ...any) { _ = "STUB: not implemented"; return }

func (c *channel) SendNoReplyInternal(method string, options ...any) {
	_ = "STUB: not implemented"
	return
}

func (c *channel) innerSendNoReply(method string, isInternal bool, options ...any) {
	_ = "STUB: not implemented"
	return
}

// ignore error actively, log only for debug

func newChannel(owner *channelOwner, object any) *channel { _ = "STUB: not implemented"; return nil }
