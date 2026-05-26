package playwright

import (
	"bufio"
	"io"
	"os"
)

type transport interface {
	Send(msg map[string]any) error
	Poll() (*message, error)
	Close() error
}

type pipeTransport struct {
	writer    io.WriteCloser
	bufReader *bufio.Reader
	closed    chan struct{}
	onClose   func() error
	process   *os.Process
}

func (t *pipeTransport) Poll() (*message, error) { _ = "STUB: not implemented"; return nil, nil }

type message struct {
	ID     int            `json:"id"`
	GUID   string         `json:"guid"`
	Method string         `json:"method,omitempty"`
	Params map[string]any `json:"params,omitempty"`
	Result map[string]any `json:"result,omitempty"`
	Error  *struct {
		Error Error `json:"error"`
	} `json:"error,omitempty"`
}

func (t *pipeTransport) Send(msg map[string]any) error { _ = "STUB: not implemented"; return nil }

func (t *pipeTransport) Close() error { _ = "STUB: not implemented"; return nil }

func (t *pipeTransport) isClosed() bool { _ = "STUB: not implemented"; return false }

func newPipeTransport(driver *PlaywrightDriver, stderr io.Writer) (transport, error) {
	_ = "STUB: not implemented"
	return *new(transport), nil
}

// playwright-cli will exit when its stdin is closed
