package playwright

type downloadImpl struct {
	page              *pageImpl
	url               string
	suggestedFilename string
	artifact          *artifactImpl
}

func (d *downloadImpl) String() string { _ = "STUB: not implemented"; return "" }

func (d *downloadImpl) Page() Page { _ = "STUB: not implemented"; return *new(Page) }

func (d *downloadImpl) URL() string { _ = "STUB: not implemented"; return "" }

func (d *downloadImpl) SuggestedFilename() string { _ = "STUB: not implemented"; return "" }

func (d *downloadImpl) Delete() error { _ = "STUB: not implemented"; return nil }

func (d *downloadImpl) Failure() error { _ = "STUB: not implemented"; return nil }

func (d *downloadImpl) Path() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *downloadImpl) SaveAs(path string) error { _ = "STUB: not implemented"; return nil }

func (d *downloadImpl) Cancel() error { _ = "STUB: not implemented"; return nil }

func newDownload(page *pageImpl, url string, suggestedFilename string, artifact *artifactImpl) *downloadImpl {
	_ = "STUB: not implemented"
	return nil
}
