package playwright

import (
	"bytes"
)

type jsHandleImpl struct {
	channelOwner
	preview string
}

func (j *jsHandleImpl) Evaluate(expression string, options ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (j *jsHandleImpl) EvaluateHandle(expression string, options ...any) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

func (j *jsHandleImpl) GetProperty(name string) (JSHandle, error) {
	_ = "STUB: not implemented"
	return *new(JSHandle), nil
}

func (j *jsHandleImpl) GetProperties() (map[string]JSHandle, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (j *jsHandleImpl) AsElement() ElementHandle {
	_ = "STUB: not implemented"
	return *new(ElementHandle)
}

func (j *jsHandleImpl) Dispose() error { _ = "STUB: not implemented"; return nil }

func (j *jsHandleImpl) String() string { _ = "STUB: not implemented"; return "" }

func (j *jsHandleImpl) JSONValue() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func parseValue(result any, refs map[float64]any) any { _ = "STUB: not implemented"; return *new(any) }

func serializeValue(value any, handles *[]*channel, depth int) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// https://github.com/golang/go/issues/2196

// had key, so convert "undefined" to "null"

func parseResult(result any) any { _ = "STUB: not implemented"; return *new(any) }

func serializeArgument(arg any) any { _ = "STUB: not implemented"; return *new(any) }

func newJSHandle(parent *channelOwner, objectType string, guid string, initializer map[string]any) *jsHandleImpl {
	_ = "STUB: not implemented"
	return nil
}

func mustBeDivisible(length int, wordSize int) int { _ = "STUB: not implemented"; return 0 }

func mustReadArray[T int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](r *bytes.Reader, v *[]T) []float64 {
	_ = "STUB: not implemented"
	return nil
}
