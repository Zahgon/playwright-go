package playwright

type apiResponseAssertionsImpl struct {
	actual APIResponse
	isNot  bool
}

func newAPIResponseAssertions(actual APIResponse, isNot bool) *apiResponseAssertionsImpl {
	_ = "STUB: not implemented"
	return nil
}

func (ar *apiResponseAssertionsImpl) Not() APIResponseAssertions {
	_ = "STUB: not implemented"
	return *new(APIResponseAssertions)
}

func (ar *apiResponseAssertionsImpl) ToBeOK() error { _ = "STUB: not implemented"; return nil }

func isTexualMimeType(mimeType string) bool { _ = "STUB: not implemented"; return false }

func subString(s string, start, length int) string { _ = "STUB: not implemented"; return "" }
