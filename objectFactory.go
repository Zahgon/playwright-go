package playwright

// dummyObject is a placeholder for unimplemented protocol objects (Android, Electron, etc.)
type dummyObject struct {
	channelOwner
}

func newDummyObject(parent *channelOwner, objectType string, guid string, initializer map[string]any) *dummyObject {
	_ = "STUB: not implemented"
	return nil
}

func createObjectFactory(parent *channelOwner, objectType string, guid string, initializer map[string]any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
