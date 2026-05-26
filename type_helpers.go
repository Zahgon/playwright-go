package playwright

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	_ = "STUB: not implemented"

	// Bool is a helper routine that allocates a new bool value
	// to store v and returns a pointer to it.
	return nil
}

func Bool(v bool) *bool {
	_ = "STUB: not implemented"

	// Int is a helper routine that allocates a new int32 value
	// to store v and returns a pointer to it.
	return nil
}

func Int(v int) *int {
	_ = "STUB: not implemented"

	// Float is a helper routine that allocates a new float64 value
	// to store v and returns a pointer to it.
	return nil
}

func Float(v float64) *float64 {
	_ = "STUB: not implemented"

	// Null will be used in certain scenarios where a strict nil pointer
	// check is not possible
	return nil
}

func Null() any {
	_ = "STUB: not implemented"

	// StringSlice is a helper routine that allocates a new StringSlice value
	// to store v and returns a pointer to it.
	return *new(any)
}

func StringSlice(v ...string) *[]string { _ = "STUB: not implemented"; return nil }

// IntSlice is a helper routine that allocates a new IntSlice value
// to store v and returns a pointer to it.
func IntSlice(v ...int) *[]int { _ = "STUB: not implemented"; return nil }

// ToOptionalStorageState converts StorageState to OptionalStorageState for use directly in [Browser.NewContext]
func (s StorageState) ToOptionalStorageState() *OptionalStorageState {
	_ = "STUB: not implemented"
	return nil
}

func (c Cookie) ToOptionalCookie() OptionalCookie {
	_ = "STUB: not implemented"
	return *new(OptionalCookie)
}
