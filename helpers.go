package playwright

import (
	"reflect"
	"regexp"
	"sync"
	"sync/atomic"

	mapset "github.com/deckarep/golang-set/v2"
)

type (
	routeHandler = func(Route)
)

func skipFieldSerialization(val reflect.Value) bool { _ = "STUB: not implemented"; return false }

func transformStructValues(in any) any { _ = "STUB: not implemented"; return *new(any) }

func transformStructIntoMapIfNeeded(inStruct any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// Merge into the base map by the JSON struct tag

// Special handling for timeout field: provide default value when nil
// This is required in Playwright v1.57+ protocol where timeout is no longer optional

// default 30s

// Skip the values when the field is a pointer (like *string) and nil.

// We use the JSON struct fields for getting the original names
// out of the field.

// Merge into the base map

// transformOptions handles the parameter data transformation
func transformOptions(options ...any) map[string]any { _ = "STUB: not implemented"; return nil }

// Case 1: No options are given

// Case 2: a single value (either struct or map) is given.

// Case 3: two values are given. The first one needs to be transformed
// to a map, the sencond one will be then get merged into the first
// base map.

// Check if the slice element type has a Timeout field and add default if so
// This is required in Playwright v1.57+ protocol where timeout is no longer optional

// default 30s

func remapValue(inMapValue reflect.Value, outStructValue reflect.Value) {
	_ = "STUB: not implemented"
	return
}

func remapMapToStruct(inputMap any, outStruct any) { _ = "STUB: not implemented"; return }

type urlMatcher struct {
	raw     any
	pattern *regexp.Regexp
	matchFn func(url string) bool
}

func newURLMatcher(urlOrPredicate any, baseURL *string, isWsUrl ...bool) *urlMatcher {
	_ = "STUB: not implemented"
	return nil
}

func (u *urlMatcher) Matches(url string) bool { _ = "STUB: not implemented"; return false }

// SameWith compares String() if urlOrPredicate is *regexp.Regexp
func (u *urlMatcher) SameWith(urlOrPredicate any) bool { _ = "STUB: not implemented"; return false }

type routeHandlerInvocation struct {
	route    Route
	complete chan bool
}

type routeHandlerEntry struct {
	matcher           *urlMatcher
	handler           routeHandler
	times             int
	count             int32
	ignoreErrors      *atomic.Bool
	activeInvocations mapset.Set[*routeHandlerInvocation]
}

func (r *routeHandlerEntry) Matches(url string) bool { _ = "STUB: not implemented"; return false }

func (r *routeHandlerEntry) Handle(route Route) chan bool { _ = "STUB: not implemented"; return nil }

// If the handler was stopped (without waiting for completion), we ignore all exceptions.

func (r *routeHandlerEntry) Stop(behavior string) {
	_ = "STUB: not implemented"
	// When a handler is manually unrouted or its page/context is closed we either
	//   - wait for the current handler invocations to finish
	//   - or do not wait, if the user opted out of it, but swallow all exceptions
	//     that happen after the unroute/close.
	return
}

func (r *routeHandlerEntry) handleInternal(route Route) chan bool {
	_ = "STUB: not implemented"
	return nil
}

func (r *routeHandlerEntry) WillExceed() bool { _ = "STUB: not implemented"; return false }

func newRouteHandlerEntry(matcher *urlMatcher, handler routeHandler, times ...int) *routeHandlerEntry {
	_ = "STUB: not implemented"
	return nil
}

func prepareInterceptionPatterns(handlers []*routeHandlerEntry) []map[string]any {
	_ = "STUB: not implemented"
	return nil
}

const defaultTimeout = 30 * 1000

type timeoutSettings struct {
	sync.RWMutex
	parent                   *timeoutSettings
	defaultTimeout           *float64
	defaultNavigationTimeout *float64
}

func (t *timeoutSettings) SetDefaultTimeout(timeout *float64) { _ = "STUB: not implemented"; return }

func (t *timeoutSettings) DefaultTimeout() *float64 { _ = "STUB: not implemented"; return nil }

func (t *timeoutSettings) Timeout(timeout ...float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t *timeoutSettings) DefaultNavigationTimeout() *float64 {
	_ = "STUB: not implemented"
	return nil
}

func (t *timeoutSettings) SetDefaultNavigationTimeout(navigationTimeout *float64) {
	_ = "STUB: not implemented"
	return
}

func (t *timeoutSettings) NavigationTimeout() float64 { _ = "STUB: not implemented"; return 0 }

func newTimeoutSettings(parent *timeoutSettings) *timeoutSettings {
	_ = "STUB: not implemented"
	return nil
}

// SelectOptionValues is the option struct for ElementHandle.Select() etc.
type SelectOptionValues struct {
	ValuesOrLabels *[]string
	Values         *[]string
	Indexes        *[]int
	Labels         *[]string
	Elements       *[]ElementHandle
}

func convertSelectOptionSet(values SelectOptionValues) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func unroute(inRoutes []*routeHandlerEntry, url any, handlers ...routeHandler) ([]*routeHandlerEntry, []*routeHandlerEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// note: compare regex expression if url is a regexp, not pointer

func serializeMapToNameAndValue(headers map[string]string) []map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// assignStructFields assigns fields from src to dest,
//
//	omitExtra determines whether to omit src's extra fields
func assignStructFields(dest, src any, omitExtra bool) error { _ = "STUB: not implemented"; return nil }

func deserializeNameAndValueToMap(headersArray []map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

type recordHarOptions struct {
	Path           string            `json:"path"`
	Content        *HarContentPolicy `json:"content,omitempty"`
	Mode           *HarMode          `json:"mode,omitempty"`
	UrlGlob        *string           `json:"urlGlob,omitempty"`
	UrlRegexSource *string           `json:"urlRegexSource,omitempty"`
	UrlRegexFlags  *string           `json:"urlRegexFlags,omitempty"`
}

type recordHarInputOptions struct {
	Path        string
	URL         any
	Mode        *HarMode
	Content     *HarContentPolicy
	OmitContent *bool
}

type harRecordingMetadata struct {
	Path    string
	Content *HarContentPolicy
}

func prepareRecordHarOptions(option recordHarInputOptions) recordHarOptions {
	_ = "STUB: not implemented"
	return *new(recordHarOptions)
}

type safeValue[T any] struct {
	sync.Mutex
	v T
}

func (s *safeValue[T]) Set(v T) { _ = "STUB: not implemented"; return }

func (s *safeValue[T]) Get() T { _ = "STUB: not implemented"; return *new(T) }
