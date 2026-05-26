package playwright

type fileChooserImpl struct {
	page          Page
	elementHandle ElementHandle
	isMultiple    bool
}

func (f *fileChooserImpl) Page() Page { _ = "STUB: not implemented"; return *new(Page) }

func (f *fileChooserImpl) Element() ElementHandle {
	_ = "STUB: not implemented"
	return *new(ElementHandle)
}

func (f *fileChooserImpl) IsMultiple() bool { _ = "STUB: not implemented"; return false }

// InputFile represents the input file for:
// - FileChooser.SetFiles()
// - ElementHandle.SetInputFiles()
// - Page.SetInputFiles()
type InputFile struct {
	Name     string `json:"name"`
	MimeType string `json:"mimeType,omitempty"`
	Buffer   []byte `json:"buffer"`
}

func (f *fileChooserImpl) SetFiles(files any, options ...FileChooserSetFilesOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func newFileChooser(page Page, elementHandle ElementHandle, isMultiple bool) *fileChooserImpl {
	_ = "STUB: not implemented"
	return nil
}
