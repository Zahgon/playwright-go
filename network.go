package playwright

type rawHeaders struct {
	headersArray []NameValue
	headersMap   map[string][]string
}

func (r *rawHeaders) Get(name string) string { _ = "STUB: not implemented"; return "" }

func (r *rawHeaders) GetAll(name string) []string { _ = "STUB: not implemented"; return nil }

func (r *rawHeaders) Headers() map[string]string { _ = "STUB: not implemented"; return nil }

func (r *rawHeaders) HeadersArray() []NameValue { _ = "STUB: not implemented"; return nil }

func newRawHeaders(headers any) *rawHeaders { _ = "STUB: not implemented"; return nil }
